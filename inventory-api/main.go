package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", homeHandler)

	// Route untuk items
	http.HandleFunc("/items", ItemsHandler)

	// Route untuk menghapus Item berdasarkan ID
	http.HandleFunc("/items/", deleteItemHandler)

	fmt.Println("Server Running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

