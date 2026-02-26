package main

import (
	"fmt"
	"net/http"
)

func main() {
	connectDatabase()

	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/items", getItems)
	http.HandleFunc("/items", itemsHandler)

	fmt.Println("Server Berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "CRUD With Database Running")
}