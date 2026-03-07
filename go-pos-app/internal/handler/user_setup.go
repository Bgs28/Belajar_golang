package handler

import (
	"database/sql"
	"go-pos-app/internal/repository"
	"go-pos-app/internal/service"
)

func NewUserHandler(db *sql.DB) *UserHandler {

	userRepo := repository.UserRepository{
		DB: db,
	}

	userService := service.UserService{
		Repo: &userRepo,
	}

	return &UserHandler{
		Service: &userService,
	}
}