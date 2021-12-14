package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	h "main/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", h.HomeHandler)
	h.HandleFuncUsers(r)
	h.HandleFuncProducts(r)
	h.HandleFuncPayments(r)
	h.HandleFuncCustomers(r)
	h.HandleFuncBaskets(r)
	h.HandleFuncBrands(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
