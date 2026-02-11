package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/geoffmilleraz/signet/internal/core/domain"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiAdapter struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewGeminiAdapter(ctx context.Context, apiKey string) (*GeminiAdapter, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	model := client.GenerativeModel("gemini-2.0-flash-exp") // Use the latest available model
	
	// Set structured output schema if supported/needed
	// For now, we'll use prompt engineering to ensure JSON output
	
	return &GeminiAdapter{
		client: client,
		model:  model,
	}, nil
}

func (a *GeminiAdapter) Analyze(ctx context.Context, diff []byte, policy domain.Policy) (domain.Evidence, error) {
	prompt := fmt.Sprintf(`You are the Overseer, a Principal SRE and Security Architect. 
Analyze the following code diff for operational reliability risks and architectural violations.
Focus on "Silent Killers":
1. Signal Handling (SIGTERM/SIGINT)
2. Resource Leaks (Connections/Goroutines)
3. Thundering Herd (Retry jitter/backoff)
4. Hardcoded Secrets
5. Schema Compatibility

Output MUST be a valid JSON object matching this structure:
{
  "verdict": "PASS | WARN | FAIL",
  "findings": [
    {
      "type": "RELIABILITY | SECURITY | ARCHITECTURE",
      "severity": "LOW | MEDIUM | HIGH",
      "file": "string",
      "line": 0,
      "message": "string",
      "remediation": "string"
    }
  ]
}

Code Diff:
%s`, string(diff))

	resp, err := a.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return domain.Evidence{}, fmt.Errorf("gemini generation failed: %w", err)
	}

	if len(resp.Candidates) == 0 {
		return domain.Evidence{}, fmt.Errorf("no candidates returned from gemini")
	}

	// Extract JSON from response (naive implementation)
	// In production, use response_mime_type: "application/json" if supported by SDK
	var analysis struct {
		Verdict  string           `json:"verdict"`
		Findings []domain.Finding `json:"findings"`
	}

	// Assuming the first part of the first candidate is the text
	part := resp.Candidates[0].Content.Parts[0]
	if text, ok := part.(genai.Text); ok {
		if err := json.Unmarshal([]byte(text), &analysis); err != nil {
			// Handle cases where LLM might wrap JSON in backticks
			return domain.Evidence{}, fmt.Errorf("failed to parse gemini response: %w", err)
		}
	}

	return domain.Evidence{
		Type:     "LLM_ANALYSIS",
		Findings: analysis.Findings,
	}, nil
}
