package main

import (
	"encoding/json"
	"net/http"
)

func getItems(w http.ResponseWriter, r *http.Request) {
	items, err := readData()
	if err != nil {
		http.Error(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	items, _ := readData()

	for _, item := range items {
		if item.ID == newItem.ID {
			// ส่ง HTTP 409 Conflict กลับไป หากมี ID ซ้ำ
			http.Error(w, "ID ซ้ำแม่เหยดด", http.StatusConflict)
			return
		}
	}

	items = append(items, newItem)

	if err := writeData(items); err != nil {
		http.Error(w, "Error saving data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	items, _ := readData()
	var newItems []Item
	found := false

	for _, item := range items {
		if item.ID != id {
			newItems = append(newItems, item)
		} else {
			found = true
		}
	}

	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	writeData(newItems)
	w.WriteHeader(http.StatusNoContent)
}
