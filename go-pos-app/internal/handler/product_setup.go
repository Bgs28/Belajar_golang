package handler

import (
	"database/sql"
	"go-pos-app/internal/repository"
	"go-pos-app/internal/service"
)

func NewProductHandler(db *sql.DB) *ProductHandler {

	productRepo := repository.ProductRepository{
		DB: db,
	}

	productService := service.ProductService{
		Repo: &productRepo,
	}

	return &ProductHandler{
		Service: &productService,
	}
}