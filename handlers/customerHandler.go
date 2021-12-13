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

// HandleFuncCustomers contains all of the HandleFunc(path, f) functions of "/customers" endpoint.
func HandleFuncCustomers(r *mux.Router) {
	r.HandleFunc("/customers", GetCustomers).Methods("GET")
	r.HandleFunc("/customers/id/{id}", GetCustomerByID).Methods("GET")
	r.HandleFunc("/customers/userID/{userID}", GetCustomerByUserID).Methods("GET")
	r.HandleFunc("/customers/purchaseAmount/{purchaseAmount}", GetCustomersByPurchaseAmount).Methods("GET")
	r.HandleFunc("/customers/orderQty/{orderQty}", GetCustomersByOrderQty).Methods("GET")
	r.HandleFunc("/customers", CreateCustomer).Methods("POST")
	r.HandleFunc("/customers", UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers", DeleteCustomer).Methods("DELETE")
}

// GetCustomers gets all the customer information from "customers.json" file and prints them.
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	fmt.Fprint(w, string(customersByte))
}

// CreateCustomer creates a customer with the given information and saves it in "customers.json" file.
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	CheckError(err)
	err = json.NewEncoder(w).Encode(&customer)
	CheckError(err)

	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	json.Unmarshal(customersByte, &customers)
	customers = append(customers, customer)
	customersByte, err = json.MarshalIndent(customers, "", "	")
	CheckError(err)

	err = os.WriteFile("data/customers.json", customersByte, 0644)
	CheckError(err)
}

// UpdateCustomer updates the customer information in "customers.json" file with the given ID.
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	CheckError(err)
	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	json.Unmarshal(customersByte, &customers)

	for i := range customers {
		if customers[i].ID == customer.ID {
			customers[i].UserID = customer.UserID
			customers[i].PurchaseAmount = customer.PurchaseAmount
			customers[i].OrderQty = customer.OrderQty
		}
	}
	customersByte, err = json.MarshalIndent(customers, "", "	")
	CheckError(err)

	err = os.WriteFile("data/customers.json", customersByte, 0644)
	CheckError(err)
}

// DeleteCustomer deletes the customer information from "customers.json" with the given ID.
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	CheckError(err)
	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	json.Unmarshal(customersByte, &customers)

	index := -1
	for i := range customers {
		if customers[i].ID == customer.ID {
			index = i
		}
	}
	customers = append(customers[:index], customers[index+1:]...)
	customersByte, err = json.MarshalIndent(customers, "", "	")
	CheckError(err)

	err = os.WriteFile("data/customers.json", customersByte, 0644)
	CheckError(err)
}

// GetCustomerByID gets the customer from "customers.json" file by ID field provided in URL.
// URL must contain "/customers/id/" followed by an ID (int) value.
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	CheckError(err)

	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	json.Unmarshal(customersByte, &customers)
	index := -1
	for i := range customers {
		if customers[i].ID == id {
			index = i
		}
	}
	customersByte, err = json.Marshal(customers[index])
	CheckError(err)

	fmt.Fprint(w, string(customersByte))
}

// GetCustomerByUserID gets the customer from "customers.json" file by UserID field provided in URL.
// URL must contain "/customers/userID/" followed by an UserID (int) value.
func GetCustomerByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	CheckError(err)

	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	json.Unmarshal(customersByte, &customers)
	index := -1
	for i := range customers {
		if customers[i].UserID == userID {
			index = i
		}
	}
	customersByte, err = json.Marshal(customers[index])
	CheckError(err)

	fmt.Fprint(w, string(customersByte))
}

// GetCustomersByPurchaseAmount gets the customer(s) from "customers.json" file by PurchaseAmount field provided in URL.
// URL must contain "/customers/price/" followed by an PurchaseAmount (float64) value.
func GetCustomersByPurchaseAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	var resCustomers []models.Customer
	params := mux.Vars(r)
	purchaseAmount, err := strconv.ParseFloat(params["purchaseAmount"], 64)
	CheckError(err)

	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	json.Unmarshal(customersByte, &customers)
	for i := range customers {
		if customers[i].PurchaseAmount == purchaseAmount {
			resCustomers = append(resCustomers, customers[i])
		}
	}
	customersByte, err = json.Marshal(resCustomers)
	CheckError(err)

	fmt.Fprint(w, string(customersByte))
}

// GetCustomersByOrderQty gets the customer(s) from "customers.json" file by OrderQty field provided in URL.
// URL must contain "/customers/orderQty/" followed by an OrderQty (int) value.
func GetCustomersByOrderQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	var resCustomers []models.Customer
	params := mux.Vars(r)
	orderQty, err := strconv.ParseFloat(params["orderQty"], 64)
	CheckError(err)

	customersByte, err := os.ReadFile("data/customers.json")
	CheckError(err)
	json.Unmarshal(customersByte, &customers)
	for i := range customers {
		if customers[i].OrderQty == int(orderQty) {
			resCustomers = append(resCustomers, customers[i])
		}
	}
	customersByte, err = json.Marshal(resCustomers)
	CheckError(err)

	fmt.Fprint(w, string(customersByte))
}
