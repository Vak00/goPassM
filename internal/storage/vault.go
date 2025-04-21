package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Vak00/goPassM/internal/model"
)

const FileName = "vault.json"

// Save all entries to a Json file
func SaveEntries(entries []model.Entry) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(FileName, data, 0600)
}

// Add an arrya of Entry to the current file
func AddEntry(service string, login string, pass string) {
	existingEntries, err := GetEntriesFromFile()

	if err != nil {
		log.Fatal(err)
		return
	}

	entry := model.Entry{Service: service, Username: login, Password: pass}
	newEntries := append(existingEntries, entry)

	errSave := SaveEntries(newEntries)

	if errSave != nil {
		log.Fatal(errSave)
	} else {
		fmt.Println("Your informations are saved ! ")
	}

}

// Load all the entry from a given Json file
func GetEntriesFromFile() ([]model.Entry, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Send an empty list if the file doesnt exist
			return []model.Entry{}, nil
		}
		return nil, err
	}

	var entries []model.Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
