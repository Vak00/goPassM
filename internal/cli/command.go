package cli

import "github.com/Vak00/goPassM/internal/service"

type Command struct {
	Name        string
	Alias       string
	Description string
	Action      func(*service.VaultService)
}
