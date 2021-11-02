package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Item struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.Method, r.URL.Path, time.Since(start)) }()
			f(rw, r)
		}
	}
}

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(rw, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func GetItems(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "get items")
}

func AddItem(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "add item")
}

func UpdateItem(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "update item")
}

func RemoveItem(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "remove item")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "index.html")
	})

	itemRouter := r.PathPrefix("/api/items").Subrouter()
	itemRouter.HandleFunc("/", Chain(GetItems, Method("GET"), Logging()))
	itemRouter.HandleFunc("/add", Chain(AddItem, Method("POST"), Logging()))
	itemRouter.HandleFunc("/update/{id}", Chain(UpdateItem, Method("PUT"), Logging()))
	itemRouter.HandleFunc("/remove/{id}", Chain(RemoveItem, Method("DELETE"), Logging()))

	http.ListenAndServe(":3000", r)
}
