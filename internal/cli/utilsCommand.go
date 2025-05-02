package cli

import (
	"fmt"
	"os"

	"github.com/Vak00/goPassM/internal/service"
)

func commandHelp(_ *service.VaultService) {
	fmt.Println("Help command here")
}

func CommandQuit(_ *service.VaultService) {
	// TODO: Do the encryption of the data here
	fmt.Println()
	fmt.Println("Bye ! 👋")
	os.Exit(0)
}
