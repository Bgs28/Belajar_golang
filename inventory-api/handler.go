package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome API is Running")
}

// ===============
// Handler Get Items
// ===============

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	default:
		http.Error(w, "Method Tidak Valid", http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item

	// Decode JSON dari request Body
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate id Sederhana
	newItem.ID = len(items) + 1

	// tambahkan ke Slice
	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// ===============
// Handler Delete Item
// ===============

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Tidak Valid", http.StatusMethodNotAllowed)
		return
	}

	// ambil ID dari Url
	idStr := r.URL.Path[len("/items/"):]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID Tidak Valid", http.StatusBadRequest)
		return
	}

	// cari hapus Item
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Item Tidak Ditemukan", http.StatusNotFound)
}