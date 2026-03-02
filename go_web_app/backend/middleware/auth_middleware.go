package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserEmailKey contextKey = "userEmail"
const TokenKey contextKey = "tokenString"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w,r, "/login", http.StatusSeeOther)
			return 
		}

		tokenString := cookie.Value

		token, err := jwt.Parse(tokenString, func(token * jwt.Token)(interface{}, error){
			const jwtKey = "my_secret_key"
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return 
		}

		email := claims["email"].(string)

		// simpan ke context
		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		ctx = context.WithValue(ctx, TokenKey, tokenString)

		next(w, r.WithContext(ctx))
	}
}