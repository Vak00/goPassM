package cli

import "github.com/Vak00/goPassM/internal/store"

type Command struct {
	Name        string
	Alias       string
	Description string
	Action      func(*store.VaultStore)
}
