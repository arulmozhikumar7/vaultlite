package cmd

import (
	"errors"
	"fmt"
	"github.com/arulmozhikumar7/vaultlite/internal/storage"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <key> <new-value>",
	Short: "Update the value of an existing secret",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		newValue := args[1]

		err := storage.UpdateSecret(key, newValue)
		if err != nil {
			if errors.Is(err, storage.ErrKeyNotFound) {
				return fmt.Errorf("update failed: key '%s' not found", key)
			}
			return fmt.Errorf("update failed: %w", err)
		}

		fmt.Println("Secret updated successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
