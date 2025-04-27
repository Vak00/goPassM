package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// Get the deried key from aes.sha256
func GetDerivedKey(password string) [32]byte {
	return sha256.Sum256([]byte(password))
}

// Encrypt based on the text and the key + a nonce (Nnumber used Once)
// The nonce is used to avoid two same messages are the same
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	// create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// create GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	// generate random number (length = aesGCM.NonceSize())
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	// encrypt
	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	// return nonce+text
	return append(nonce, ciphertext...), nil
}

// Decrypt a giver text based on the key and nonce
// ciphertext: contains the nonce+text
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	// create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// create GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// split the cipher and the nonce
	// format: nonce+text
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("❌ Error occured during the decryption: ciphertext too short")
	}
	nonce, actualCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// decrypt
	plaintext, err := aesGCM.Open(nil, nonce, actualCiphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
