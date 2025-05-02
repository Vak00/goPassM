package storage

import (
	"os"
)

const FileName = ".data"

// // Save all entries to a Json file
// func SaveEntries(entries []model.Entry) error {
// 	data, err := json.MarshalIndent(entries, "", "  ")
// 	if err != nil {
// 		return err
// 	}
// 	return os.WriteFile(FileName, data, 0600)
// }

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
