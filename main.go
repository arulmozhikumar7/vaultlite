package main

import (
	"fmt"
	"os"

	"github.com/arulmozhikumar7/vaultlite/cmd"
	"github.com/arulmozhikumar7/vaultlite/internal/config"
)

func main() {
	// Show version and exit
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			fmt.Println("VaultLite", cmd.Version)
			return
		}
	}

	// Skip config check ONLY for `init`
	if len(os.Args) >= 2 && os.Args[1] != "init" {
		valid, err := config.ConfigExistsAndValid()
		if err != nil {
			fmt.Println("Error validating config:", err)
			os.Exit(1)
		}
		if !valid {
			fmt.Println("Vault is not initialized. Please run: vaultlite init")
			os.Exit(1)
		}
	}


	cmd.Execute()
}
