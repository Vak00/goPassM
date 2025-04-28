package cli

import (
	"fmt"
	"os"
)

// Allow the user to add a new Entry
func commandAdd() {
	fmt.Println("✏️  Add new entry")
	fmt.Println()
	entryService, _, _ := askForOneEntry()

	fmt.Println()
	fmt.Println("✅ New entry for " + entryService + " saved !")
}

func commandHelp() {
	fmt.Println("Help command here")
}

func commandList() {
	fmt.Println("List command here")
}

func CommandQuit() {
	// TODO: Do the encryption of the data here
	fmt.Println()
	fmt.Println("Bye ! 👋")
	os.Exit(0)
}
