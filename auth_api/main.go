package main

import (
	"auth_api/config"
	"auth_api/middleware"
	"log"
	"net/http"

	"auth_api/handlers"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.Handle("/profile", middleware.JWTAuth(http.HandlerFunc(handlers.Profile)))

	log.Println("Server Running on : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}