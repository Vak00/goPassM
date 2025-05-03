package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/Vak00/goPassM/internal/storage"
	"golang.org/x/crypto/pbkdf2"
)

// Generate a random salt of 16 bytes
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to generate random salt: %w", err)
	}
	return salt, nil
}

// Get the deried key from aes.sha256 (32 bytes)
func getDerivedKeyFromMasterHash(hashedPassword string, salt []byte) ([]byte, error) {
	iterations := 100000

	key := pbkdf2.Key([]byte(hashedPassword), salt, iterations, 32, sha256.New)
	return key, nil
}

// Encrypt based on the text and the key + a nonce (Nnumber used Once)
// The nonce is used to avoid two same messages are the same
//
// return: byte (nonce+ciphertext) concatenated
func encryptData(plaintext []byte, key []byte) (nonce []byte, ciphertext []byte, err error) {
	// create cipher block based on the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, fmt.Errorf("❌ Failed to create cipher: %v", err)
	}
	// create GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, fmt.Errorf("❌ Failed to create GCM: %v", err)
	}
	// generate the nonce (length = aesGCM.NonceSize())
	nonce = make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, nil, fmt.Errorf("❌ Failed to created the nonce")
	}
	ciphertext = aesGCM.Seal(nil, nonce, plaintext, nil)
	return nonce, ciphertext, nil
}

// Public method to encrypted and save into a file all the data
func SaveEncryptedVault(plaintext []byte, masterHashPass string) error {
	salt, errSalt := generateSalt()
	if errSalt != nil {
		return errSalt
	}

	key, errkey := getDerivedKeyFromMasterHash(masterHashPass, salt)
	if errkey != nil {
		return errkey
	}

	nonce, ciphertext, err := encryptData(plaintext, key)
	if err != nil {
		return nil
	}
	return storage.SaveToFile(salt, nonce, ciphertext)
}

// Load and decrypted the data from the vault file
func LoadEncryptedData(masterHashPass string) ([]byte, error) {
	encryptedData, errFile := storage.GetEncryptedFileContent()
	if errFile != nil {
		return nil, fmt.Errorf("❌ Failed to load the encrypted file: %v", errFile)
	}

	// If there is no file
	if len(encryptedData) == 0 {
		return nil, fmt.Errorf("⚠️ Current vault file not found")
	}

	// Extract the salt(16 bytes), the nonce(16 bytes) and the ciphertext
	salt := encryptedData[:16]
	nonce := encryptedData[16:28] // 12 bytes de nonce (typique GCM)
	cipherText := encryptedData[28:]

	// Get the key
	key, errKey := getDerivedKeyFromMasterHash(masterHashPass, salt)
	if errKey != nil {
		return nil, fmt.Errorf("❌ Failed to generate the key: %v", errKey)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to create cipher: %v", err)
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to create GCM: %v", err)
	}

	// Decrypt
	plaintext, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to decrypt: %v", err)
	}

	return plaintext, nil
}
