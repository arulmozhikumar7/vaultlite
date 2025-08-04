package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/arulmozhikumar7/vaultlite/internal/storage"
)

var showMeta bool


var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Retrieve a secret by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value, err := storage.GetSecret(key, showMeta)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVar(&showMeta, "show-meta", false, "Show metadata for the secret")
}
