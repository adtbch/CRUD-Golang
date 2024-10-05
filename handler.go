package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "modernc.org/sqlite"
)

// CreateProductHandler untuk membuat produk baru
func CreateProductHandler(c echo.Context) error {
	productDTO := new(ProductDTO)
	if err := c.Bind(productDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stmt, err := db.Prepare("INSERT INTO products(name, price) VALUES(?, ?)")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res, err := stmt.Exec(productDTO.Name, productDTO.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	id, _ := res.LastInsertId()
	product := &Product{
		ID:    int(id),
		Name:  productDTO.Name,
		Price: productDTO.Price,
	}

	return c.JSON(http.StatusCreated, product)
}

// GetProductsHandler untuk mengambil semua produk
func GetProductsHandler(c echo.Context) error {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}

// GetProductByIDHandler untuk mengambil produk berdasarkan ID
func GetProductByIDHandler(c echo.Context) error {
	id := c.Param("id")
	var product Product
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, product)
}

// UpdateProductHandler untuk mengubah data produk
func UpdateProductHandler(c echo.Context) error {
	id := c.Param("id")
	productDTO := new(ProductDTO)
	if err := c.Bind(productDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	_, err = stmt.Exec(productDTO.Name, productDTO.Price, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	product := &Product{
		ID:    0, // ID tidak diubah
		Name:  productDTO.Name,
		Price: productDTO.Price,
	}

	return c.JSON(http.StatusOK, product)
}

// DeleteProductHandler untuk menghapus produk
func DeleteProductHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Product deleted successfully")
}
