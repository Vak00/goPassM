package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/Vak00/goPassM/internal/cli"
	"github.com/Vak00/goPassM/internal/crypto"
)

const masterFilePath = ".master"

// Return true if the master file exists, else false
func IsMasterFilePresent() bool {
	if _, err := os.Stat(masterFilePath); err == nil {
		return true
	} else {
		return false
	}
}

// Get the content of the master file
func getHashFromMasterFile() (string, error) {
	data, err := os.ReadFile(masterFilePath)
	if err != nil {
		return "", err
	}
	return string(strings.TrimSpace(string(data))), nil
}

func saveMasterHash(base65password string) {
	err := os.WriteFile(masterFilePath, []byte(base65password), 0644)
	if err != nil {
		fmt.Println("❌ Error occured when trying to save the master password, abort : " + err.Error())
		os.Exit(1)
	}
	fmt.Println("✅ Master password saved !")
}

func AskForPasswordCreation() {
	fmt.Println("You have to enter a master password in order to register some entries.")
	fmt.Println("This password will be requested each time you run this app. Try not forget it ! 😎")

	passwordOne := cli.AskPassword("Enter your master password : ")
	passwordTwo := cli.AskPassword("Confirm your master password : ")

	if strings.Compare(passwordOne, passwordTwo) != 0 {
		fmt.Println("👎 Password are not the same, focus !")
		os.Exit(1)
	}
	hash, err := crypto.HashString(passwordOne)
	if err != nil {
		fmt.Println("❌ Error during the creation of the hash : " + err.Error())
		os.Exit(1)
	}
	saveMasterHash(hash)
}

// Ask and compare the user master password
func AskForMasterPassword() {
	userpassword := cli.AskPassword("Enter the master password : ")

	hashFromFile, err := getHashFromMasterFile()
	fmt.Println("HASH FROM FILE: " + hashFromFile)
	if err != nil {
		fmt.Println("Error occured when try to get the master password from file: " + err.Error())
		os.Exit(1)
	}

	isSame := crypto.IsSameHash(hashFromFile, userpassword)
	fmt.Printf("IS SMAE : %t", isSame)
	if !crypto.IsSameHash(hashFromFile, userpassword) {
		fmt.Println("⛔️ Password incorrect, exit")
		os.Exit(0)
	}
	fmt.Println("✅ Access granted")
}
