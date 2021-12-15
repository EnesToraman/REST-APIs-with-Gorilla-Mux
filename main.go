package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	h "main/handlers"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static")))
	h.HandleFuncUsers(r)
	h.HandleFuncProducts(r)
	h.HandleFuncPayments(r)
	h.HandleFuncCustomers(r)
	h.HandleFuncBaskets(r)
	h.HandleFuncBrands(r)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
