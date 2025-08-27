package main

import (
	"fmt"
	"net/http"
)

var users = map[string]string{
	"user1": "password123",
	"admin": "adminpass",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			nama := r.FormValue("nama")
			email := r.FormValue("email")

			fmt.Fprintf(w, "Nama Anda: %s, Email: %s", nama, email)
		} else {
			http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

		fmt.Println("Server berjalan di http://localhost:8080")
		http.ListenAndServe(":8080", nil)
}