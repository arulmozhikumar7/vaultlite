package cmd

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/arulmozhikumar7/vaultlite/internal/config"
	"github.com/arulmozhikumar7/vaultlite/internal/cipher"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize VaultLite with encryption key and IV",
	Run: func(cmd *cobra.Command, args []string) {
		ok, err := config.ConfigExistsAndValid()
		if err != nil {
			fmt.Println("Failed to read config:", err)
			os.Exit(1)
		}
		if ok {
			fmt.Println("Vault is already initialized. Skipping.")
			return
		}


		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter passphrase: ")
		pass1, _ := reader.ReadString('\n')
		pass1 = strings.TrimSpace(pass1)

		fmt.Print("Confirm passphrase: ")
		pass2, _ := reader.ReadString('\n')
		pass2 = strings.TrimSpace(pass2)

		if pass1 != pass2 {
			fmt.Println("Passphrases do not match. Aborting.")
			os.Exit(1)
		}

		// Generate random 16-byte salt
		salt := make([]byte, 16)
		_, err = rand.Read(salt)
		if err != nil {
			fmt.Println("Failed to generate salt:", err)
			os.Exit(1)
		}

		key, iv := encryption.DeriveKeyAndIV(pass1, salt)

		cfg := config.Config{
			Key:  base64.StdEncoding.EncodeToString(key),
			IV:   base64.StdEncoding.EncodeToString(iv),
			Salt: base64.StdEncoding.EncodeToString(salt),
		}

		err = config.SaveConfig(cfg)
		if err != nil {
			fmt.Println("Failed to save config:", err)
			os.Exit(1)
		}

		fmt.Println("Vault initialized successfully.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
