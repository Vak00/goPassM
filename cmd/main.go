package main

import (
	"github.com/Vak00/goPassM/internal/auth"
	"github.com/Vak00/goPassM/internal/cli"
	"github.com/Vak00/goPassM/internal/signals"
)

func main() {
	// Start signals listening
	signals.StartSignalListener()

	if auth.IsMasterFilePresent() {
		auth.AskForMasterPassword()
	} else {
		auth.AskForPasswordCreation()
	}

	cli.AskAndShowMenu()
}
