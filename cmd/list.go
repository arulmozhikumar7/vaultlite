package cmd

import (
	"fmt"

	"github.com/arulmozhikumar7/vaultlite/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys",
	Run: func(cmd *cobra.Command, args []string) {
		keys, err := storage.ListSecrets()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Secrets:")
		for _, key := range keys {
			fmt.Println(" -", key)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
