package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tonpcst/go-microservice-prisma-postgresql/database"
	"github.com/tonpcst/go-microservice-prisma-postgresql/prisma/db"
)

// @Summary     Create a new user
// @Description Adds a new user to the database
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       body  body  map[string]string  true  "User Data"
// @Success     201  {object}  map[string]interface{}
// @Router      /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient

	var paylod struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&paylod)
	if err != nil || paylod.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input"))
		return
	}

	newUser, err := pClient.Client.User.CreateOne(db.User.Name.Set(paylod.Name)).Exec(pClient.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	out, _ := json.Marshal(newUser)
	w.Write(out)
}

// @Summary     Get all users
// @Description Retrieves all users from the database
// @Tags        users
// @Produce     json
// @Success     200 {object} map[string]interface{}
// @Router      /users/all [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient
	allUsers, err := pClient.Client.User.FindMany().Exec(pClient.Context)
	if err != nil {
		fmt.Println("Cannot fetch users")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	usersMap := map[string]interface{}{"users": allUsers}
	out, _ := json.MarshalIndent(usersMap, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// @Summary     Get a user by ID
// @Description Retrieves a user by ID from the database
// @Tags        users
// @Produce     json
// @Param       id   path      int  true  "User ID"
// @Success     200  {object}  map[string]interface{}
// @Router      /users/byId/{id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is required"))
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	user, err := pClient.Client.User.FindUnique(
		db.User.ID.Equals(userID),
	).Exec(pClient.Context)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	userMap := map[string]interface{}{"user": user}
	out, _ := json.MarshalIndent(userMap, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// @Summary Update a user
// @Description Updates a user's information in the database
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body map[string]string true "User Data"
// @Success 200 {object} map[string]interface{}
// @Router /users/update/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is required"))
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	var paylod struct {
		Name string `json:"name"`
	}
	err = json.NewDecoder(r.Body).Decode(&paylod)
	if err != nil || paylod.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input"))
		return
	}

	updatedUser, err := pClient.Client.User.FindUnique(db.User.ID.Equals(userID)).Update(db.User.Name.Set(paylod.Name)).Exec(pClient.Context)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Failed to update user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	out, _ := json.Marshal(updatedUser)
	w.Write(out)
}

// @Summary Delete a user
// @Description Deletes a user from the database
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/delete/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is required"))
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found"))
		return
	}

	_, err = pClient.Client.User.FindUnique(db.User.ID.Equals(userID)).Delete().Exec(pClient.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to delete user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
