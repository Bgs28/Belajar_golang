package main

import (
	"fmt"
	"go-pos-app/internal/routes"
	"go-pos-app/pkg/database"
	"net/http"
)

func main() {

	db, err := database.ConnectDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "POS System Running")
	})

	routes.RegisterRoutes(db)

	fmt.Println("Server running on: 8080")

	http.ListenAndServe(":8080", nil)
}