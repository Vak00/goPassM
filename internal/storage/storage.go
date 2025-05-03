package storage

import (
	"encoding/json"
	"os"

	"github.com/Vak00/goPassM/internal/store"
)

const FileName = ".data"

// Save the data in this order: salt(16 bytes) + nonce(32 bytes) + ciphertext concatenated
func SaveToFile(salt, nonce, ciphertext []byte) error {
	data := append(salt, nonce...)
	data = append(data, ciphertext...)
	return os.WriteFile(FileName, data, 0600)
}

// Get the content of the main encryted file
func GetEncryptedFileContent() ([]byte, error) {
	fileData, err := os.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Send an empty list if the file doesnt exist
			return []byte{}, nil
		}
		return nil, err
	}
	return fileData, nil
}

// Create the json file based on the vault content and return it
func CreateJsonFileFromVault(vault *store.VaultStore) ([]byte, error) {
	data, err := json.MarshalIndent(vault.GetAll(), "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
