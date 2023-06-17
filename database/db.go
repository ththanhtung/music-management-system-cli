package database

import (
	"encoding/json"
	"os"
)

// for simplicity i will use json as database

func ReadFromDB[T any](filename string) (map[string]*T, error) {
	// Load existing JSON data from file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[string]*T, 0)
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func SaveToDB[T any](filename string, item T) error {
	// Write data to file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(item)
	if err != nil {
		return err
	}

	return nil
}
