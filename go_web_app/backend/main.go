package main

import (
	"go_web_app/config"
	"go_web_app/handlers"
	"go_web_app/middleware"

	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "../frontend/register.html")
			return
		}

		if r.Method == http.MethodPost {
			handlers.Register(w,r)
			return
		}
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet{
			http.ServeFile(w,r, "../frontend/login.html")
			return
		}

		if r.Method == http.MethodPost {
			handlers.Login(w,r)
			return
		}
	})
	http.Handle("/profile", middleware.AuthMiddleware(handlers.Profile))
	http.HandleFunc("/logout", handlers.Logout)

	log.Println("Server Running on : 8080")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(http.DefaultServeMux)))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", "GET, POST, OPTIONS")
		
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

