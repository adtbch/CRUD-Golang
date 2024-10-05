package main

// ProductDTO digunakan untuk menerima input dari client
type ProductDTO struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
}
