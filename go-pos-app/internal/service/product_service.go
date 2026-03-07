package service

import (
	"go-pos-app/internal/model"
	"go-pos-app/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) CreateProduct(product model.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetProduct() ([]model.Product, error) {
	return s.Repo.GetProduct()
}

func (s *ProductService) GetProductByID(id int) (model.Product, error){
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(product model.Product) error {
	return s.Repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error{
	return s.Repo.DeleteProduct(id)
}