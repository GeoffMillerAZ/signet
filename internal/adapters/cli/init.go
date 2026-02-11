package cli

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a repository for Signet governance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing Signet structure...")
		
		dirs := []string{".signet/core", ".signet/app"}
		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
				return
			}
		}

		// Write default policy to .signet/core/policy.cue
		defaultPolicy := `// Signet Global Policy
#GlobalPolicy: {
    steps: {
        scan: {
            required: true
            timeout: "10m"
        }
    }
}`
		if err := os.WriteFile(".signet/core/policy.cue", []byte(defaultPolicy), 0644); err != nil {
			fmt.Printf("Error writing policy file: %v\n", err)
			return
		}

		fmt.Println("Successfully initialized .signet/ directory.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
