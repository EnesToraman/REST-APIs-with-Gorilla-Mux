package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"main/models"
)

// HandleFuncProducts contains all of the HandleFunc(path, f) functions of "/products" endpoint.
func HandleFuncProducts(r *mux.Router) {
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/sku/{sku}", GetProductBySKU).Methods("GET")
	r.HandleFunc("/products/name/{name}", GetProductByName).Methods("GET")
	r.HandleFunc("/products/price/{price}", GetProductsByPrice).Methods("GET")
	r.HandleFunc("/products/stock/{stock}", GetProductsByStock).Methods("GET")
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products", UpdateProduct).Methods("PUT")
	r.HandleFunc("/products", DeleteProduct).Methods("DELETE")
}

// GetProducts gets all the product information from "products.json" file and prints them.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	fmt.Fprint(w, string(productsBytes))
}

// CreateProduct creates a product with the given information and saves it in "products.json" file.
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	err = json.NewEncoder(w).Encode(&product)
	CheckError(err)

	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	json.Unmarshal(productsBytes, &products)
	products = append(products, product)
	productsByte, err := json.MarshalIndent(products, "", "	")
	CheckError(err)

	err = os.WriteFile("data/products.json", productsByte, 0644)
	CheckError(err)
}

// UpdateProduct updates the product information in "products.json" file with the given ID.
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	json.Unmarshal(productsBytes, &products)

	for i := range products {
		if products[i].ID == product.ID {
			products[i].SKU = product.SKU
			products[i].Name = product.Name
			products[i].Price = product.Price
			products[i].Stock = product.Stock
		}
	}
	productsByte, err := json.MarshalIndent(products, "", "	")
	CheckError(err)

	err = os.WriteFile("data/products.json", productsByte, 0644)
	CheckError(err)
}

// DeleteProduct deletes the product information from "products.json" with the given ID.
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	json.Unmarshal(productsBytes, &products)

	index := -1
	for i := range products {
		if products[i].ID == product.ID {
			index = i
		}
	}
	products = append(products[:index], products[index+1:]...)
	productsByte, err := json.MarshalIndent(products, "", "	")
	CheckError(err)

	err = os.WriteFile("data/products.json", productsByte, 0644)
	CheckError(err)
}

// GetProductBySKU gets the product from "products.json" file by SKU field provided in URL.
// URL must contain "/products/sku/" followed by an SKU (string) value.
func GetProductBySKU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	params := mux.Vars(r)
	sku := params["sku"]

	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	json.Unmarshal(productsBytes, &products)
	index := -1
	for i := range products {
		if products[i].SKU == sku {
			index = i
		}
	}
	productByte, err := json.Marshal(products[index])
	CheckError(err)

	fmt.Fprint(w, string(productByte))
}

// GetProductByName gets the product from "products.json" file by Name field provided in URL.
// URL must contain "/products/name/" followed by an Name (string) value.
func GetProductByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	params := mux.Vars(r)
	name := params["name"]

	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	json.Unmarshal(productsBytes, &products)
	index := -1
	for i := range products {
		if products[i].Name == name {
			index = i
		}
	}
	productByte, err := json.Marshal(products[index])
	CheckError(err)

	fmt.Fprint(w, string(productByte))
}

// GetProductsByPrice gets the product from "products.json" file by Price field provided in URL.
// URL must contain "/products/price/" followed by an Price (float64) value.
func GetProductsByPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	var resProducts []models.Product
	params := mux.Vars(r)
	price, err := strconv.ParseFloat(params["price"], 64)
	CheckError(err)

	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	json.Unmarshal(productsBytes, &products)
	for i := range products {
		if products[i].Price == price {
			resProducts = append(resProducts, products[i])
		}
	}
	productByte, err := json.Marshal(resProducts)
	CheckError(err)

	fmt.Fprint(w, string(productByte))
}

// GetProductsByStock gets the product from "products.json" file by Stock field provided in URL.
// URL must contain "/products/Stock/" followed by an Stock (int) value.
func GetProductsByStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	var resProducts []models.Product
	params := mux.Vars(r)
	stock, err := strconv.Atoi(params["stock"])
	CheckError(err)

	productsBytes, err := os.ReadFile("data/products.json")
	CheckError(err)
	json.Unmarshal(productsBytes, &products)
	for i := range products {
		if products[i].Stock == stock {
			resProducts = append(resProducts, products[i])
		}
	}
	productByte, err := json.Marshal(resProducts)
	CheckError(err)

	fmt.Fprint(w, string(productByte))
}
