package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello")
	})

	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/{id}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprintf(rw, "Got id %s \n", id)
	}).Methods("GET")

	http.ListenAndServe(":3000", r)
}
