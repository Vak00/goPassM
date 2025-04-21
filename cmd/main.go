package main

import (
	"fmt"
	"log"

	"github.com/Vak00/goPassM/internal/storage"
)

func main() {

	// entry := model.Entry{
	// 	Service:  "gmail",
	// 	Username: "johnDoe@gmail.com",
	// 	Password: "hehehehehehehehehe",
	// }

	const vaultFile = "vault.json"

	entries, err := storage.GetEntriesFromFile(vaultFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-- Current entries -- ")

	for _, entry := range entries {
		fmt.Println("passaword saved for the serice: " + entry.Service)
	}

	// entries = append(entries, entry)

	// err = storage.SaveEntries(entries, vaultFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Password saved bro ! ")
}
