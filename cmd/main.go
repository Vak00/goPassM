package main

import (
	"github.com/Vak00/goPassM/internal/auth"
	"github.com/Vak00/goPassM/internal/cli"
)

func main() {

	if auth.IsMasterFilePresent() {
		auth.AskForMasterPassword()
	} else {
		auth.AskForPasswordCreation()
	}

	cli.Run()
}
