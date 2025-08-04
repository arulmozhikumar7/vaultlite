package cmd

import (
	"fmt"

	"github.com/arulmozhikumar7/vaultlite/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <key> <value>",
	Short: "Add a new secret",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		err := storage.AddSecret(key, value)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Secret added successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
