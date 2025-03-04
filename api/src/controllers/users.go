package controllers

import (
	database "api/src/dataBase"
	"api/src/models"
	"api/src/repositores"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser is a function that creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(req, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repositore := repositores.NewUserRepository(db)

	userBD, err := repositore.CreateUser(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusCreated, userBD)

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
