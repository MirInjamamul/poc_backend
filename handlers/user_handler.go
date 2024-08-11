package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mirinjamamul/go-poc-api/database"
	"github.com/Mirinjamamul/go-poc-api/models"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.CreateUser(user)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := database.GetUsers()
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, found := database.GetUser(params["id"])
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.UpdateUser(params["id"], user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	database.DeleteUser(params["id"])
	w.WriteHeader(http.StatusNoContent)
}
