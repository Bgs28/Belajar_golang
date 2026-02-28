package handlers

import (
	"auth_api/config"
	"auth_api/models"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtkey = []byte("my_secret_key")

func Profile(w http.ResponseWriter, r *http.Request){
	userID := r.Context().Value("user_id")
	email := r.Context().Value("email")

	response := map[string]interface{}{
		"message": "Welcome to Your Profile",
		"user_id": userID,
		"email": email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedUser models.User

	err = config.DB.QueryRow(
		"SELECT id, email, password_user FROM users WHERE email = ?",
		user.Email,
	).Scan(&storedUser.ID, &storedUser.Email, &storedUser.Password)

	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(storedUser.Password),
		[]byte(user.Password),
	)

	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"email": storedUser.Email,
	})

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func Register(w http.ResponseWriter, r *http.Request){
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = config.DB.Exec(
		"INSERT INTO users (email, password_user) VALUES (?, ?)",
		user.Email,
		string(hashedPassword),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User Registered Succesfully"))
}

