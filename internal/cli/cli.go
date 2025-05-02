package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/Vak00/goPassM/internal/service"
	"golang.org/x/term"
)

var commands = []Command{
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
		Name:        "help",
		Alias:       "h",
		Description: "Show help menu",
		Action:      commandHelp,
	},
	{
		Name:        "quit",
		Alias:       "q",
		Description: "Quit the app",
		Action:      CommandQuit,
	},
}

func AskAndShowMenu(vault *service.VaultService) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n -- Menu -- ")
	for _, cmd := range commands {
		fmt.Printf(" - %s (%s): %s\n", cmd.Name, cmd.Alias, cmd.Description)
	}

	fmt.Println()
	fmt.Print("Choose an option : ")
	commandInput, _ := reader.ReadString('\n')
	commandCleared := clear(commandInput)

	for _, cmd := range commands {
		if cmd.Name == commandCleared || cmd.Alias == commandCleared {
			cmd.Action(vault)
			AskAndShowMenu(vault)
			return
		}
	}

	fmt.Printf("❌ Error, the command '%s' is not available.\n", commandCleared)
}

// Delete end of line delimiter
func clear(s string) string {
	return strings.TrimSpace(s)
}

// Run the form to return entry fields
func AskForOneEntry() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📖 Enter your service name : ")
	service, _ := reader.ReadString('\n')

	fmt.Print("👤 Enter your login : ")
	login, _ := reader.ReadString('\n')

	password := AskPassword("Enter your password : ")

	return clear(service), clear(login), clear(password)
}

// Ask for a password to the user
func AskPassword(prompt string) string {
	fmt.Print("🔑 " + prompt)
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	// To force the next Print() on the next line
	fmt.Println()
	return string(bytePassword)
}
