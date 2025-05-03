package service

import (
	"fmt"

	"github.com/Vak00/goPassM/internal/crypto"
	"github.com/Vak00/goPassM/internal/input"
	"github.com/Vak00/goPassM/internal/model"
	"github.com/Vak00/goPassM/internal/storage"
	"github.com/Vak00/goPassM/internal/store"
)

func AddEntry(vault *store.VaultStore) {
	fmt.Println()
	fmt.Println("Add entry service here")

	login, service, pass := input.AskForOneEntry()

	vault.AddEntry(model.Entry{Service: service, Username: login, Password: pass})
	fmt.Println("✅ Entrée sauvegardée")
}

func ListEntry(vault *store.VaultStore) {
	entries := vault.GetAll()

	if len(entries) == 0 {
		fmt.Println("📭 No entries in the vault.")
		return
	}

	fmt.Println("\n📒 Vault entries:")
	fmt.Println("────────────────────────────────────────────")
	fmt.Printf(" %-20s | %-20s\n", "🔖 Service", "👤 Username")
	fmt.Println("────────────────────────────────────────────")

	for _, entry := range entries {
		fmt.Printf(" %-20s | %-20s\n", entry.Service, entry.Username)
	}
	fmt.Println("────────────────────────────────────────────")
}

func EditEntry(vault *store.VaultStore) {
	fmt.Println("Edit entry service here")
}

func DeleteEntry(vault *store.VaultStore) {
	fmt.Println("Delete entry service here")
}

func SaveVault(vault *store.VaultStore) {
	data, err := storage.CreateJsonFileFromVault(vault)
	if err != nil {
		fmt.Println("❌ Impossible to create the vault file : " + err.Error())
		return
	}
	crypto.SaveEncryptedVault(data, "aaa")
	fmt.Println("✅ Data saved ! ")
	input.ExitApp()
}
