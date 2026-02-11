package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
	mode    string
)

var rootCmd = &cobra.Command{
	Use:   "signet",
	Short: "Signet is an autonomous governance platform",
	Long: `Signet ensures your artifacts are semantically and syntactically valid 
before promotion, using Cuelang and LLM-powered analysis.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .signet/config.cue)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "local", "runtime mode: local or connected")
}
