package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/geoffmilleraz/signet/internal/adapters/crypto"
	"github.com/geoffmilleraz/signet/internal/adapters/git"
	"github.com/geoffmilleraz/signet/internal/adapters/ledger"
	"github.com/geoffmilleraz/signet/internal/core/services"
	"github.com/spf13/cobra"
)

var (
	targetEnv string
	sealID    string
)

var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promotes a sealed artifact to a target environment",
	Run: func(cmd *cobra.Command, args []string) {
		if sealID == "" || targetEnv == "" {
			fmt.Println("Error: --seal-id and --env are required")
			os.Exit(1)
		}

		fmt.Printf("Promoting Seal %s to %s...
", sealID, targetEnv)

		// Initialize Adapters
		// In a real CLI, we'd use the production SQLite file
		ledgerAdapter, err := ledger.NewSQLiteEventStoreAdapter("signet.db")
		if err != nil {
			fmt.Printf("Error opening ledger: %v
", err)
			os.Exit(1)
		}
		cryptoAdapter := crypto.NewSHA256Adapter()
		gitAdapter := git.NewFileAdapter()

		// Initialize Service
		svc := services.NewPromotionService(ledgerAdapter, gitAdapter)

		// Execute
		prURL, err := svc.Promote(context.Background(), sealID, targetEnv)
		if err != nil {
			fmt.Printf("Promotion failed: %v
", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully created promotion PR: %s
", prURL)
		fmt.Printf("Seal ID: %s has been stamped on the deployment.
", sealID)
	},
}

func init() {
	promoteCmd.Flags().StringVar(&sealID, "seal-id", "", "The ID of the cryptographic seal to promote")
	promoteCmd.Flags().StringVar(&targetEnv, "env", "", "The target environment (e.g., prod, staging)")
	rootCmd.AddCommand(promoteCmd)
}
