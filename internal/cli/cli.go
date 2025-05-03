package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Vak00/goPassM/internal/input"
	"github.com/Vak00/goPassM/internal/service"
	"github.com/Vak00/goPassM/internal/store"
)

var commandsList = []Command{
	{
		Name:        "list",
		Alias:       "l",
		Description: "List all entries from the vault",
		Action:      service.ListEntry,
	},
	{
		Name:        "add",
		Alias:       "a",
		Description: "Add a new entry to the vault",
		Action:      service.AddEntry,
	},
	{
		Name:        "edit",
		Alias:       "e",
		Description: "Edit one entry",
		Action:      service.EditEntry,
	},
	{
		Name:        "quit",
		Alias:       "q",
		Description: "Quit the app",
		Action:      service.SaveVault,
	},
}

func AskAndShowMenu(vault *store.VaultStore) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n -- Menu -- ")
	for _, cmd := range commandsList {
		fmt.Printf(" - %s (%s): %s\n", cmd.Name, cmd.Alias, cmd.Description)
	}

	fmt.Println()
	fmt.Print("Choose an option : ")
	commandInput, _ := reader.ReadString('\n')
	commandCleared := input.Clear(commandInput)

	for _, cmd := range commandsList {
		if cmd.Name == commandCleared || cmd.Alias == commandCleared {
			cmd.Action(vault)
			AskAndShowMenu(vault)
			return
		}
	}
	fmt.Printf("❌ Error, the command '%s' is not available.\n", commandCleared)
}
