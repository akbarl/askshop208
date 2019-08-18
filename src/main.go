package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ArticlesCategoryHandler)
	r.HandleFunc("/products", ArticlesCategoryHandler)
	r.HandleFunc("/articles", ArticlesCategoryHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8181", r)
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
