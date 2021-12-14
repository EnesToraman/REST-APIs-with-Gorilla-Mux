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

// HandleFuncBrands contains all of the HandleFunc(path, f) functions of "/brands" endpoint.
func HandleFuncBrands(r *mux.Router) {
	r.HandleFunc("/brands", GetBrands).Methods("GET")
	r.HandleFunc("/brands/id/{id}", GetBrandByID).Methods("GET")
	r.HandleFunc("/brands/name/{name}", GetBrandByName).Methods("GET")
	r.HandleFunc("/brands/productQty/{productQty}", GetBrandsByProductQty).Methods("GET")
	r.HandleFunc("/brands/totalWorth/{totalWorth}", GetBrandsByTotalWorth).Methods("GET")
	r.HandleFunc("/brands", CreateBrand).Methods("POST")
	r.HandleFunc("/brands", UpdateBrand).Methods("PUT")
	r.HandleFunc("/brands", DeleteBrand).Methods("DELETE")
}

// GetBrands gets all the brand information from "brands.json" file and prints them.
func GetBrands(w http.ResponseWriter, r *http.Request) {
	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	fmt.Fprint(w, string(brandsByte))
}

// CreateBrand creates a brand with the given information and saves it in "brands.json" file.
func CreateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brands []models.Brand
	var brand models.Brand

	err := json.NewDecoder(r.Body).Decode(&brand)
	CheckError(err)
	err = json.NewEncoder(w).Encode(&brand)
	CheckError(err)

	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	json.Unmarshal(brandsByte, &brands)
	brands = append(brands, brand)
	brandsByte, err = json.MarshalIndent(brands, "", "	")
	CheckError(err)

	err = os.WriteFile("data/brands.json", brandsByte, 0644)
	CheckError(err)
}

// UpdateBrand updates the brand information in "brands.json" file with the given ID.
func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brands []models.Brand
	var brand models.Brand

	err := json.NewDecoder(r.Body).Decode(&brand)
	CheckError(err)
	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	json.Unmarshal(brandsByte, &brands)

	for i := range brands {
		if brands[i].ID == brand.ID {
			brands[i].Name = brand.Name
			brands[i].ProductQty = brand.ProductQty
			brands[i].TotalWorth = brand.TotalWorth
		}
	}
	brandsByte, err = json.MarshalIndent(brands, "", "	")
	CheckError(err)

	err = os.WriteFile("data/brands.json", brandsByte, 0644)
	CheckError(err)
}

// DeleteBrand deletes the brand information from "brands.json" with the given ID.
func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brands []models.Brand
	var brand models.Brand

	err := json.NewDecoder(r.Body).Decode(&brand)
	CheckError(err)
	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	json.Unmarshal(brandsByte, &brands)

	index := -1
	for i := range brands {
		if brands[i].ID == brand.ID {
			index = i
		}
	}
	brands = append(brands[:index], brands[index+1:]...)
	brandsByte, err = json.MarshalIndent(brands, "", "	")
	CheckError(err)

	err = os.WriteFile("data/brands.json", brandsByte, 0644)
	CheckError(err)
}

// GetBrandByID gets the brand from "brands.json" file by ID field provided in URL.
// URL must contain "/brands/id/" followed by an ID (int) value.
func GetBrandByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brands []models.Brand
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	CheckError(err)

	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	json.Unmarshal(brandsByte, &brands)
	index := -1
	for i := range brands {
		if brands[i].ID == id {
			index = i
		}
	}
	brandsByte, err = json.Marshal(brands[index])
	CheckError(err)

	fmt.Fprint(w, string(brandsByte))
}

// GetBrandByName gets the brand from "brands.json" file by Name field provided in URL.
// URL must contain "/brands/name/" followed by an Name (string) value.
func GetBrandByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brands []models.Brand
	params := mux.Vars(r)
	name := params["name"]

	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	json.Unmarshal(brandsByte, &brands)
	index := -1
	for i := range brands {
		if brands[i].Name == name {
			index = i
		}
	}
	brandsByte, err = json.Marshal(brands[index])
	CheckError(err)

	fmt.Fprint(w, string(brandsByte))
}

// GetBrandsByProductQty gets the brand(s) from "brands.json" file by Username field provided in URL.
// URL must contain "/brands/productQty/" followed by an ProductQty (int) value.
func GetBrandsByProductQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brands []models.Brand
	var resBrands []models.Brand
	params := mux.Vars(r)
	productQty, err := strconv.Atoi(params["productQty"])
	CheckError(err)

	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	json.Unmarshal(brandsByte, &brands)
	for i := range brands {
		if brands[i].ProductQty == productQty {
			resBrands = append(resBrands, brands[i])
		}
	}
	brandsByte, err = json.Marshal(resBrands)
	CheckError(err)

	fmt.Fprint(w, string(brandsByte))
}

// GetBrandsByTotalWorth gets the brand(s) from "brands.json" file by TotalWorth field provided in URL.
// URL must contain "/brands/totalWorth/" followed by an TotalWorth (float64) value.
func GetBrandsByTotalWorth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brands []models.Brand
	var resBrands []models.Brand
	params := mux.Vars(r)
	totalWorth, err := strconv.ParseFloat(params["totalWorth"], 64)
	CheckError(err)

	brandsByte, err := os.ReadFile("data/brands.json")
	CheckError(err)
	json.Unmarshal(brandsByte, &brands)
	for i := range brands {
		if brands[i].TotalWorth == totalWorth {
			resBrands = append(resBrands, brands[i])
		}
	}
	resbrandsByte, err := json.Marshal(resBrands)
	CheckError(err)

	fmt.Fprint(w, string(resbrandsByte))
}
