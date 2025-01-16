package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User Struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User
var nextID = 1

// Get All Users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get Single User
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// Create User
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// Update User
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, user := range users {
		if user.ID == id {
			users[index].Name = "Updated Name"
			json.NewEncoder(w).Encode(users[index])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// Delete User
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	// Define Routes
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Start Server
	http.ListenAndServe(":8080", r)
}
