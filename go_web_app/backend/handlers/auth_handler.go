package handlers

import (
	"html/template"
	"net/http"
	"time"

	"go_web_app/config"
	"go_web_app/middleware"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

func Logout(w http.ResponseWriter, r *http.Request){

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path: "/",
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Profile(w http.ResponseWriter, r *http.Request){
	email := r.Context().Value(middleware.UserEmailKey).(string)
	token := r.Context().Value(middleware.TokenKey).(string)

	tmpl, err := template.ParseFiles("../frontend/profile.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"Email": email,
		"Token": token,
	}

	tmpl.Execute(w, data)
}

func Login(w http.ResponseWriter, r *http.Request){
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email dan Password Wajib Diisi", http.StatusBadRequest)
		return
	}

	var storedPass string

	err := config.DB.QueryRow(
		"SELECT password_user FROM users WHERE email = ?",
		email,
	).Scan(&storedPass)

	if err != nil {
		http.Error(w, "Email Tidak ditemukan", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(storedPass),
		[]byte(password),
	)

	if err != nil {
		http.Error(w, "Password Salah", http.StatusUnauthorized)
		return
	}

	// buat token
	claims := jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Gagal Membuat Token", http.StatusInternalServerError)
		return
	}

	// simpan token ke cookie

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Path: "/",
	})

	http.Redirect(w,r, "/profile", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request){
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email dan Password Wajib Diisi", http.StatusBadRequest)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = config.DB.Exec(
		"INSERT INTO users (email, password_user) VALUES (?, ?)",
		email,
		string(hashedPass),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

http.Redirect(w,r, "/register", http.StatusSeeOther)
}

