package storage

import (
	"encoding/json"
	"os"

	"github.com/Vak00/goPassM/internal/model"
)

// Save all entries to a Json file
func SaveEntries(entries []model.Entry, filename string) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0600)
}

// Load all the entry from a given Json file
func GetEntriesFromFile(filename string) ([]model.Entry, error) {
	data, err := os.ReadFile(filename)
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
