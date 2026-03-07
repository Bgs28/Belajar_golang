package repository

import (
	"database/sql"
	"go-pos-app/internal/model"
)

type ProductRepository struct {
	DB *sql.DB
}

// update Product
func (r *ProductRepository) UpdateProduct(product model.Product) error {
	query := "UPDATE products SET name=?, price=?, stock=? WHERE id=?"

	_, err := r.DB.Exec(query, product.Name, product.Price, product.Stock, product.ID)

	return err
}

// create product
func (r *ProductRepository) CreateProduct(product model.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES (?, ?, ?)"

	_, err := r.DB.Exec(query, product.Name, product.Price, product.Stock)

	return err
}

// get Product (untuk menampilkan semua product)
func (r *ProductRepository) GetProduct() ([]model.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price, stock FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var p model.Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

// get Prodcut by ID
func (r *ProductRepository) GetProductByID(id int) (model.Product, error){
	query := "SELECT id, name, price, stock FROM products WHERE id=?"

	var product model.Product

	err := r.DB.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
	)

	return product, err
}

// delete Product
func (r *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id=?"

	_, err := r.DB.Exec(query, id)

	return err
}