package model

type Command struct {
	Name        string
	Alias       string
	Description string
	Action      func()
}
