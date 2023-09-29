package endpoints

import (
	"encoding/json"
	"fmt"
	"os"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func LoadAlbumsFromFile(filename string) ([]Album, error) {

	data, err := os.ReadFile(fmt.Sprintf("data/%s", filename))
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var albums []Album

	err = json.Unmarshal(data, &albums)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal albums: %w", err)
	}

	return albums, nil
}

func SaveAlbumsToFile(filename string, albums []Album) error {

	data, err := json.MarshalIndent(albums, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal albums: %w", err)
	}

	err = os.WriteFile(fmt.Sprintf("data/%s", filename), data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write albums to file: %w", err)
	}

	return nil
}
