package controllers

import (
	database "api/src/dataBase"
	"api/src/models"
	"api/src/repositores"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser is a function that creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err = json.Unmarshal(req, &user); err != nil {
		log.Fatal(err)
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repositore := repositores.NewUserRepository(db)

	userId, err := repositore.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created with id %d", userId.ID)))

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
