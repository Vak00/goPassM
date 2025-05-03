package store

import "github.com/Vak00/goPassM/internal/model"

type VaultStore struct {
	entries []model.Entry
}

// Init a new VaultService struct with a list of Entry
func NewVaultStore(entries []model.Entry) *VaultStore {
	return &VaultStore{entries: entries}
}

// Add a new entry to the vault
func (vault *VaultStore) AddEntry(entry model.Entry) {
	vault.entries = append(vault.entries, entry)
}

// Get all entries registered in the vault
func (vault *VaultStore) GetAll() []model.Entry {
	return vault.entries
}
