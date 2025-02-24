package controllers

import (
	"net/http"
)

// CreateUser is a function that creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) { 
	w.Write([]byte("Creating a user"))
}

// GetUsers is a function that gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}

// GetUser is a function that gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting a user"))
}

// UpdateUser is a function that updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a user"))
}

// DeleteUser is a function that deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user"))
}
