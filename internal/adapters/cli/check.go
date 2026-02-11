package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run a local governance scan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running local governance check...")
		// TODO: Implement check logic
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
