package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/geoffmilleraz/signet/internal/adapters/crypto"
	"github.com/geoffmilleraz/signet/internal/adapters/llm"
	"github.com/geoffmilleraz/signet/internal/adapters/policy"
	"github.com/geoffmilleraz/signet/internal/core/services"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run a local governance scan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running local governance check...")

		// Initialize Adapters
		cueAdapter := policy.NewCUEAdapter()
		cryptoAdapter := crypto.NewSHA256Adapter()
		llmAdapter := llm.NewMockLLMAdapter()

		// Initialize Service
		svc := services.NewValidationService(cueAdapter, llmAdapter, cryptoAdapter)

		// Mock Data for "Current Changes"
		globalPolicy := []byte(`#GlobalPolicy: { steps: { scan: { required: true } } }`)
		userPatch := []byte(`#UserPatch: { meta: { name: "cli-run" } }`)
		diff := []byte("func main() { // No SIGTERM handling }")

		// Execute
		verdict, findings, err := svc.ValidateDiff(context.Background(), globalPolicy, userPatch, diff)
		if err != nil {
			fmt.Printf("Error running check: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Verdict: %s\n", verdict)
		for _, f := range findings {
			fmt.Printf("- [%s] %s: %s\n", f.Severity, f.Type, f.Message)
		}

		if verdict == "FAIL" {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
