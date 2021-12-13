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

// HandleFuncPayments contains all of the HandleFunc(path, f) functions of "/payments" endpoint.
func HandleFuncPayments(r *mux.Router) {
	r.HandleFunc("/payments", GetPayments).Methods("GET")
	r.HandleFunc("/payments/userID/{userID}", GetPaymentsByUserID).Methods("GET")
	r.HandleFunc("/payments/amount/{amount}", GetPaymentsByAmount).Methods("GET")
	r.HandleFunc("/payments/discount/{discount}", GetPaymentsByDiscount).Methods("GET")
	r.HandleFunc("/payments/tax/{tax}", GetPaymentsByTax).Methods("GET")
	r.HandleFunc("/payments", CreatePayment).Methods("POST")
	r.HandleFunc("/payments", UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments", DeletePayment).Methods("DELETE")
}

// GetPayments gets all the payment information from "payments.json" file and prints them.
func GetPayments(w http.ResponseWriter, r *http.Request) {
	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	fmt.Fprint(w, string(paymentsByte))
}

// CreatePayment creates a payment with the given information and saves it in "payments.json" file.
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []models.Payment
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	CheckError(err)
	err = json.NewEncoder(w).Encode(&payment)
	CheckError(err)

	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	json.Unmarshal(paymentsByte, &payments)
	payments = append(payments, payment)
	paymentsByte, err = json.MarshalIndent(payments, "", "	")
	CheckError(err)

	err = os.WriteFile("data/payments.json", paymentsByte, 0644)
	CheckError(err)
}

// UpdatePayment updates the payment information in "payments.json" file with the given ID.
func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []models.Payment
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	CheckError(err)
	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	json.Unmarshal(paymentsByte, &payments)

	for i := range payments {
		if payments[i].ID == payment.ID {
			payments[i].UserID = payment.UserID
			payments[i].Amount = payment.Amount
			payments[i].Discount = payment.Discount
			payments[i].Tax = payment.Tax
		}
	}
	paymentsByte, err = json.MarshalIndent(payments, "", "	")
	CheckError(err)

	err = os.WriteFile("data/payments.json", paymentsByte, 0644)
	CheckError(err)
}

// DeletePayment deletes the payment information from "payments.json" with the given ID.
func DeletePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []models.Payment
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	CheckError(err)
	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	json.Unmarshal(paymentsByte, &payments)

	index := -1
	for i := range payments {
		if payments[i].ID == payment.ID {
			index = i
		}
	}
	payments = append(payments[:index], payments[index+1:]...)
	paymentsByte, err = json.MarshalIndent(payments, "", "	")
	CheckError(err)

	err = os.WriteFile("data/payments.json", paymentsByte, 0644)
	CheckError(err)
}

// GetPaymentsByUserID gets the payment from "payments.json" file by UserID field provided in URL.
// URL must contain "/payments/userID/" followed by an UserID (int) value.
func GetPaymentsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []models.Payment
	var resPayments []models.Payment
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	CheckError(err)

	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	json.Unmarshal(paymentsByte, &payments)
	for i := range payments {
		if payments[i].UserID == userID {
			resPayments = append(resPayments, payments[i])
		}
	}
	paymentByte, err := json.Marshal(resPayments)
	CheckError(err)

	fmt.Fprint(w, string(paymentByte))
}

// GetPaymentsByAmount gets the payment from "payments.json" file by Amount field provided in URL.
// URL must contain "/payments/amount/" followed by an Amount (float64) value.
func GetPaymentsByAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []models.Payment
	var resPayments []models.Payment
	params := mux.Vars(r)
	amount, err := strconv.ParseFloat(params["amount"], 64)
	CheckError(err)

	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	json.Unmarshal(paymentsByte, &payments)
	for i := range payments {
		if payments[i].Amount == amount {
			resPayments = append(resPayments, payments[i])
		}
	}
	paymentByte, err := json.Marshal(resPayments)
	CheckError(err)

	fmt.Fprint(w, string(paymentByte))
}

// GetPaymentsByDiscount gets the payment from "payments.json" file by Discount field provided in URL.
// URL must contain "/payments/discount/" followed by an Discount (float64) value.
func GetPaymentsByDiscount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []models.Payment
	var resPayments []models.Payment
	params := mux.Vars(r)
	discount, err := strconv.ParseFloat(params["discount"], 64)
	CheckError(err)

	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	json.Unmarshal(paymentsByte, &payments)
	for i := range payments {
		if payments[i].Discount == discount {
			resPayments = append(resPayments, payments[i])
		}
	}
	paymentByte, err := json.Marshal(resPayments)
	CheckError(err)

	fmt.Fprint(w, string(paymentByte))
}

// GetPaymentsByTax gets the payment from "payments.json" file by Tax field provided in URL.
// URL must contain "/payments/tax/" followed by an Tax (float64) value.
func GetPaymentsByTax(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []models.Payment
	var resPayments []models.Payment
	params := mux.Vars(r)
	tax, err := strconv.ParseFloat(params["tax"], 64)
	CheckError(err)

	paymentsByte, err := os.ReadFile("data/payments.json")
	CheckError(err)
	json.Unmarshal(paymentsByte, &payments)
	for i := range payments {
		if payments[i].Tax == tax {
			resPayments = append(resPayments, payments[i])
		}
	}
	paymentByte, err := json.Marshal(resPayments)
	CheckError(err)

	fmt.Fprint(w, string(paymentByte))
}
