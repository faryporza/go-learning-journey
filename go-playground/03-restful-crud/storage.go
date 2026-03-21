package main

import (
	"encoding/json"
	"os"
	"sync"
)

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

const dataFile = "data.json"

var mu sync.Mutex

func readData() ([]Item, error) {
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func writeData(items []Item) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}
