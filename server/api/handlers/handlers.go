package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/kovarniykrab/servermain/database"
	"gitlab.com/kovarniykrab/servermain/domain"

	"github.com/gorilla/mux"
)

type userHandler struct {
	domain *domain.Model
}

func (h *userHandler) GetAllUsersHandler(res http.ResponseWriter, req *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(users)
}

func (h *userHandler) GetUser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(req, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserByID(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(user)
}

func (h *userHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var user domain.Model

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	userCreated, err := database.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(userCreated)
}

func (h *userHandler) UpdateUser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user domain.Model
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(req, err.Error(), http.StatusBadRequest)
		return
	}

	updateUser, err := database.UpdateUser(id, user)
}
