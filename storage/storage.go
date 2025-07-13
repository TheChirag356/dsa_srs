package storage

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func LoadCards(filename string) (*CardStore, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty store
			return &CardStore{}, nil
		}
		return nil, err
	}

	var store CardStore
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, err
	}
	return &store, nil
}

func SaveCards(filename string, store *CardStore) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal("Error creating dir:", err)
	}
	data, err := json.MarshalIndent(store, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
