package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// for simplicity i will use json as database

func ReadMapFromDB[T any](filename string) (map[string]*T, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a map of string to MusicTrackItem
	var tracks map[string]*T
	err = json.Unmarshal(data, &tracks)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func SaveMapToDB[T any](filename string, item T) error {
	// Write data to file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
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

func ReadSliceFromDB[T any](filename string) ([]*T, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of *MusicLibraryItem
	var items []*T
	err = json.Unmarshal(data, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func SaveSliceToDB[T any](filename string, items []*T) error {
	// Convert the data to JSON
	jsonData, err := json.Marshal(items)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
