package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(rw http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(rw, "%s %s is %d years old", user.FirstName, user.LastName, user.Age)
	})

	http.HandleFunc("/encode", func(rw http.ResponseWriter, r *http.Request) {
		arifin := User{
			FirstName: "arifin",
			LastName:  "tahu",
			Age:       24,
		}
		json.NewEncoder(rw).Encode(arifin)
	})

	http.ListenAndServe(":8080", nil)
}
