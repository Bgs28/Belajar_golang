package main

// ================
// Model
// ================

type Item struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Stock int `json:"stock"`
}

// ================
// Storage (Sementara pakai slice)
// ================
var items = []Item{
	{ID: 1, Name: "Laptop", Stock: 10},
	{ID: 2, Name: "Samsung", Stock: 35},
}
