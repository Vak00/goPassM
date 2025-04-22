package crypto

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// Return the hash in base64 of the given string
func HashString(text string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hashBytes), nil
}

// Return true if the given hash is the same than the given string, else fasle
func IsSameHash(base64Hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(base64Hash), []byte(password))
	return err == nil
}
