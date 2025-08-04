package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "v1.0.0"
var showVersion bool

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "VaultLite is a simple CLI secrets manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Print version and exit if --version is used
		if showVersion {
			fmt.Println("VaultLite", Version)
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show version and exit")
}
