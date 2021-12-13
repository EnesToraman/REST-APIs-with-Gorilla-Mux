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

// HandleFuncBaskets contains all of the HandleFunc(path, f) functions of "/baskets" endpoint.
func HandleFuncBaskets(r *mux.Router) {
	r.HandleFunc("/baskets", GetBaskets).Methods("GET")
	r.HandleFunc("/baskets/userID/{userID}", GetBasketByUserID).Methods("GET")
	r.HandleFunc("/baskets/productID/{productID}", GetBasketsByProductID).Methods("GET")
	r.HandleFunc("/baskets/sku/{sku}", GetBasketsBySKU).Methods("GET")
	r.HandleFunc("/baskets/quantity/{quantity}", GetBasketsByQuantity).Methods("GET")
	r.HandleFunc("/baskets", CreateBasket).Methods("POST")
	r.HandleFunc("/baskets", UpdateBasket).Methods("PUT")
	r.HandleFunc("/baskets", DeleteBasket).Methods("DELETE")
}

// GetBaskets gets all the basket information from "baskets.json" file and prints them.
func GetBaskets(w http.ResponseWriter, r *http.Request) {
	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	fmt.Fprint(w, string(basketsByte))
}

// CreateBasket creates a basket with the given information and saves it in "baskets.json" file.
func CreateBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var baskets []models.Basket
	var basket models.Basket

	err := json.NewDecoder(r.Body).Decode(&basket)
	CheckError(err)
	err = json.NewEncoder(w).Encode(&basket)
	CheckError(err)

	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	json.Unmarshal(basketsByte, &baskets)
	baskets = append(baskets, basket)
	basketsByte, err = json.MarshalIndent(baskets, "", "	")
	CheckError(err)

	err = os.WriteFile("data/baskets.json", basketsByte, 0644)
	CheckError(err)
}

// UpdateBasket updates the basket information in "baskets.json" file with the given ID.
func UpdateBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var baskets []models.Basket
	var basket models.Basket

	err := json.NewDecoder(r.Body).Decode(&basket)
	CheckError(err)
	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	json.Unmarshal(basketsByte, &baskets)

	for i := range baskets {
		for j := range baskets[i].Products {
			if baskets[i].Products[j].ProductID == basket.Products[j].ProductID {
				baskets[i].Products[j].SKU = basket.Products[j].SKU
				baskets[i].Products[j].Quantity = basket.Products[j].Quantity
			}
		}
	}
	basketsByte, err = json.MarshalIndent(baskets, "", "	")
	CheckError(err)

	err = os.WriteFile("data/baskets.json", basketsByte, 0644)
	CheckError(err)
}

// DeleteBasket deletes the basket information from "baskets.json" with the given userID.
func DeleteBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var baskets []models.Basket
	var basket models.Basket

	err := json.NewDecoder(r.Body).Decode(&basket)
	CheckError(err)
	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	json.Unmarshal(basketsByte, &baskets)

	index := -1
	for i := range baskets {
		if baskets[i].UserID == basket.UserID {
			index = i
		}
	}
	baskets = append(baskets[:index], baskets[index+1:]...)
	basketsByte, err = json.MarshalIndent(baskets, "", "	")
	CheckError(err)

	err = os.WriteFile("data/baskets.json", basketsByte, 0644)
	CheckError(err)
}

// GetBasketByUserID gets the basket from "baskets.json" file by UserID field provided in URL.
// URL must contain "/baskets/userID/" followed by an UserID (int) value.
func GetBasketByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var baskets []models.Basket
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	CheckError(err)

	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	json.Unmarshal(basketsByte, &baskets)
	index := -1
	for i := range baskets {
		if baskets[i].UserID == userID {
			index = i
		}
	}
	basketsByte, err = json.Marshal(baskets[index])
	CheckError(err)

	fmt.Fprint(w, string(basketsByte))
}

// GetBasketsByProductID gets the basket(s) from "baskets.json" file by ProductID field provided in URL.
// URL must contain "/baskets/productID/" followed by an ProductID (int) value.
func GetBasketsByProductID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var baskets []models.Basket
	var resBaskets []models.Basket
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["productID"])
	CheckError(err)

	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	json.Unmarshal(basketsByte, &baskets)
	for i := range baskets {
		for j := range baskets[i].Products {
			if baskets[i].Products[j].ProductID == productID {
				resBaskets = append(resBaskets, baskets[i])
			}
		}
	}
	basketsByte, err = json.Marshal(resBaskets)
	CheckError(err)

	fmt.Fprint(w, string(basketsByte))
}

// GetBasketsBySKU gets the basket(s) from "baskets.json" file by SKU field provided in URL.
// URL must contain "/baskets/sku/" followed by an SKU (string) value.
func GetBasketsBySKU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var baskets []models.Basket
	var resBaskets []models.Basket
	params := mux.Vars(r)
	sku := params["sku"]

	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	json.Unmarshal(basketsByte, &baskets)
	for i := range baskets {
		for j := range baskets[i].Products {
			if baskets[i].Products[j].SKU == sku {
				resBaskets = append(resBaskets, baskets[i])
			}
		}
	}
	basketsByte, err = json.Marshal(resBaskets)
	CheckError(err)

	fmt.Fprint(w, string(basketsByte))
}

// GetBasketsByQuantity gets the basket(s) from "baskets.json" file by Quantity field provided in URL.
// URL must contain "/baskets/quantity/" followed by an Quantity (int) value.
func GetBasketsByQuantity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var baskets []models.Basket
	var resBaskets []models.Basket
	params := mux.Vars(r)
	quantity, err := strconv.Atoi(params["quantity"])
	CheckError(err)

	basketsByte, err := os.ReadFile("data/baskets.json")
	CheckError(err)
	json.Unmarshal(basketsByte, &baskets)
	for i := range baskets {
		for j := range baskets[i].Products {
			if baskets[i].Products[j].Quantity == quantity {
				resBaskets = append(resBaskets, baskets[i])
			}
		}
	}
	basketsByte, err = json.Marshal(resBaskets)
	CheckError(err)

	fmt.Fprint(w, string(basketsByte))
}
