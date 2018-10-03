package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes set the routes for the web service.
func Routes() {
	http.HandleFunc("/user", GetUser)
}

// GetUser returns corresponding user details.
func GetUser(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "george",
		Email: "george@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&u)
}
