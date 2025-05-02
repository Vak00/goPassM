package cli

import (
	"fmt"
	"os"

	"github.com/Vak00/goPassM/internal/service"
)

// Allow the user to add a new Entry
func commandAdd() {
	fmt.Println("✏️  Add new entry")
	fmt.Println()
	login, ser, password := askForOneEntry()

	service.AddEntry(login, ser, password)
}

func commandHelp() {
	fmt.Println("Help command here")
}

func commandList() {
	fmt.Println("List command here")
	service.ListEntry()
}

func commandEdit() {
	fmt.Println("Edit command here")
	service.EditEntry()
}

func CommandQuit() {
	// TODO: Do the encryption of the data here
	fmt.Println()
	fmt.Println("Bye ! 👋")
	os.Exit(0)
}
