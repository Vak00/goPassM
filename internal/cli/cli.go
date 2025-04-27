package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/Vak00/goPassM/internal/crypto"
	"github.com/Vak00/goPassM/internal/storage"
	"golang.org/x/term"
)

func Run() {
	if len(os.Args) == 1 {
		fmt.Println("You need to provide at least one argument.")
		return
	}

	firstArg := os.Args[1]
	if firstArg != "add" {
		fmt.Println("First arg has to be 'add'")
		return
	}

	// Get the input from the user now
	service, login, password := askForOneEntry()

	storage.AddEntry(service, login, password)
}

// Delete end of line delimiter
func clear(stringToClean string) string {
	return strings.Replace(stringToClean, "\n", "", -1)
}

// Run the form to return entry fields
func askForOneEntry() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📖 Enter your service name : ")
	service, _ := reader.ReadString('\n')

	fmt.Print("👤 Enter your login : ")
	login, _ := reader.ReadString('\n')

	password := AskPassword("Enter your password : ")
	hashedPassword, _ := crypto.HashString(clear(password))

	return clear(service), clear(login), hashedPassword
}

// Ask for a password to the user
func AskPassword(prompt string) string {
	fmt.Print("🔑 " + prompt)
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	// To force the next Print() on the next line
	fmt.Println()
	return string(bytePassword)
}
