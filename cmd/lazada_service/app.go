package main

import (
	"net/http"

	"github.com/akbarl/askshop208/internal/service/lazada_service/lazada_handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//lazada_handlers.NewAuthorizationHandler()
	r.HandleFunc("/auth", lazada_handlers.AuthenticationHandler)
	r.HandleFunc("/products", lazada_handlers.ProductsHandler)
	http.ListenAndServe(":8181", r)
}
