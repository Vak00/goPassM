package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Vak00/goPassM/internal/storage"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	firstArg := os.Args[1]

	if firstArg != "add" {
		fmt.Println("You need to provide at least one argument.")
		return
	}

	// Get the input from the user now
	fmt.Println("Enter your service name : ")
	service, _ := reader.ReadString('\n')

	fmt.Println("Enter your login : ")
	login, _ := reader.ReadString('\n')

	fmt.Println("Enter your password : ")
	password, _ := reader.ReadString('\n')

	storage.AddEntry(clear(service), clear(login), clear(password))
}

// Delete end of line delimiter
func clear(stringToClean string) string {
	return strings.Replace(stringToClean, "\n", "", -1)
}
