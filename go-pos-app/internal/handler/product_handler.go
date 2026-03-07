package handler

import (
	"encoding/json"
	"go-pos-app/internal/model"
	"go-pos-app/internal/service"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Service *service.ProductService
}

func (h *ProductHandler) HandleProduct (w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodPost:
		h.CreateProduct(w,r)

	case http.MethodGet:
		if r.URL.Query().Get("id") != "" {
			h.GetProductByID(w,r)
		} else {
			h.GetProduct(w,r)
		}
	case http.MethodPut:
		h.UpdateProduct(w,r)

	case http.MethodDelete:
		h.DeleteProduct(w,r)

	default:
		http.Error(w, "Methode Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateProduct(product)

	if err != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product Created",
	})
}


func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request){

	products, err := h.Service.GetProduct()

	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	if h.Service ==  nil {
		log.Println("Service Masih Nil")
	}

	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request){
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	product, err := h.Service.GetProductByID(id)

	if err != nil {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func(h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request){
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	var product model.Product

	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "invalid Body", http.StatusBadRequest)
		return
	}

	product.ID = id

	err = h.Service.UpdateProduct(product)

	if err != nil {
		http.Error(w, "Failed Update Product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"Message": "Product Updated",
	})
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request){
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteProduct(id)

	if err != nil {
		http.Error(w, "Failed Delete Product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product Deleted",
	})
}