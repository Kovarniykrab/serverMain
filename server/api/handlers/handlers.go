package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/kovarniykrab/servermain/server/database"
	"gitlab.com/kovarniykrab/servermain/server/domain"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	db database.Database
}

func NewUserHandler(db database.Database) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetAllUsersHandler(res http.ResponseWriter, req *http.Request) {
	users, err := h.db.GetAllUsers()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(users)
}

func (h *UserHandler) GetUser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.db.GetUserByID(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(user)
}

func (h *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var user domain.Model

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	userCreated, err := h.db.CreateUser(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(userCreated)
}

func (h *UserHandler) UpdateUser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user domain.Model
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	updateUser, err := h.db.UpdateUser(id, user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(updateUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.db.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
