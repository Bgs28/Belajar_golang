package handler

import (
	"encoding/json"
	"go-pos-app/internal/model"
	"go-pos-app/internal/service"
	"log"
	"net/http"
)
 
type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUsers(w,r)
	case http.MethodPost:
		h.CreateUsers(w,r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request){
	users, err := h.Service.GetUsers()

	if err != nil {
		log.Println("Eror Get User:", err)
		http.Error(w, "Failed fetch Users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUsers(w http.ResponseWriter, r *http.Request){
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateUsers(user)

	if err != nil {
		http.Error(w, "Failed create User", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Created",
	})
}