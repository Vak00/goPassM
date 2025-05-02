package service

import "github.com/Vak00/goPassM/internal/model"

type VaultService struct {
	entries []model.Entry
}

// Init a new VaultService struct with a list of Entry
func NewVaultService(entries []model.Entry) *VaultService {
	return &VaultService{entries: entries}
}

// Add a new entry to the vault
func (vault *VaultService) AddEntry(entry model.Entry) {
	vault.entries = append(vault.entries, entry)
}

// Get all entries registered in the vault
func (vault *VaultService) GetAll() []model.Entry {
	return vault.entries
}
