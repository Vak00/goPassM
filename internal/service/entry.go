package service

import "fmt"

func AddEntry(login string, service string, password string) {
	fmt.Println()
	fmt.Println("Add entry service here")

	fmt.Println("Recievd : " + login + " service : " + service + " pass : " + password)
}

func ListEntry() {
	fmt.Println("List entry service here")
}

func EditEntry() {
	fmt.Println("Edit entry service here")
}

func DeleteEntry() {
	fmt.Println("Delete entry service here")
}
