package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"simple-http-server/models"
)

var Users = []models.User{
	{
		ID:    1,
		Name:  "Budi",
		Email: "budi@example.com",
	},
	{
		ID:    2,
		Name:  "Ani",
		Email: "ani@example.com",
	},
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		idParam := r.URL.Query().Get("id")

		if idParam == "" {
			json.NewEncoder(w).Encode(Users)
			return
		}

		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		for _, user := range Users {
			if user.ID == id {
				json.NewEncoder(w).Encode(user)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		var newUser models.User

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		newUser.ID = len(Users) + 1
		Users = append(Users, newUser)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
		return
	}

	if r.Method == http.MethodDelete {
		idParam := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		for i, user := range Users {
			if user.ID == id {
				Users = append(Users[:i], Users[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPut {
		idParam := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var updatedUser models.User

		err = json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		for i, user := range Users {
			if user.ID == id {
				Users[i].Name = updatedUser.Name
				Users[i].Email = updatedUser.Email

				json.NewEncoder(w).Encode(Users[i])
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}