package main

import (
	"github.com/labstack/echo/v4"
)

// RegisterRoutes digunakan untuk mendaftarkan routing ke handler CRUD
func RegisterRoutes(e *echo.Echo) {
	e.POST("/products", CreateProductHandler)
	e.GET("/products", GetProductsHandler)
	e.GET("/products/:id", GetProductByIDHandler)
	e.PUT("/products/:id", UpdateProductHandler)
	e.DELETE("/products/:id", DeleteProductHandler)
}
