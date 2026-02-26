package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func itemsHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	case http.MethodPut:
		updateItem(w, r)
	case http.MethodDelete:
		deleteItem(w, r)
	default:
		http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter, r *http.Request){
	rows, err := db.Query("SELECT id, name_items, stock FROM items")

	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"INSERT INTO items (name_items, stock) VALUES (?, ?)",
		item.Name,
		item.Stock,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	item.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is Required", http.StatusBadRequest)
		return
	}

	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"UPDATE items SET name_items = ?, stock = ? WHERE id = ?",
		item.Name,
		item.Stock,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	item.ID, _ = strconv.Atoi(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"delete from items where id = ? ", 
		id,
	)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("item deleted Succesfully"))
}