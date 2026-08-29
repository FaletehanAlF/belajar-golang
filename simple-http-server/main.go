package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"simple-http-server/handlers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Welcome to my Go HTTP Server!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is my first Go HTTP Server project.")
}

func helloAPIHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello from Go API!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/api/hello", helloAPIHandler)
	http.HandleFunc("/api/users", handlers.UsersHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}