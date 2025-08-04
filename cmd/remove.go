package cmd

import (
	"fmt"

	"github.com/arulmozhikumar7/vaultlite/internal/storage"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <key>",
	Short: "Remove a secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		err := storage.RemoveSecret(key)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Secret removed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
