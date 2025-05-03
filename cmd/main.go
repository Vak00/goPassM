package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Vak00/goPassM/internal/auth"
	"github.com/Vak00/goPassM/internal/cli"
	"github.com/Vak00/goPassM/internal/crypto"
	"github.com/Vak00/goPassM/internal/model"
	"github.com/Vak00/goPassM/internal/signals"
	"github.com/Vak00/goPassM/internal/store"
)

func main() {
	// Start signals listening
	signals.StartSignalListener()

	// Authentification
	var masterPassword string
	if auth.IsMasterFilePresent() {
		masterPassword = auth.AskForMasterPassword()
	} else {
		masterPassword = auth.AskForPasswordCreation()
	}

	// Load exsting entries if exists, else empty slice
	var entries []model.Entry
	plainText, err := crypto.LoadEncryptedData(masterPassword)
	if err != nil {
		fmt.Println("⚠️ Impossible to decrypt the vault, starting with an empty one. " + err.Error())
		fmt.Println()
		entries = []model.Entry{}
	} else {
		err := json.Unmarshal(plainText, &entries)
		if err != nil {
			fmt.Println("❌ Failed to load the Json : " + err.Error())
			os.Exit(1)
		}
	}

	vaultService := store.NewVaultStore(entries)
	cli.AskAndShowMenu(vaultService)
}
