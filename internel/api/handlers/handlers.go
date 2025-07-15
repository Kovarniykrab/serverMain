package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/kovarniykrab/servermain/internel/domain"
)

// SearchUser godoc
// @Summary     поиск пользователя
// @Description поиск всех польхователей
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.UserForm  true  "JSON"
// @Success      200  {object} domain.UserForm
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      409  {object}  Error
// @Failure      500  {object}  Error
// @Router      / [POST].
func SearchUser(app *App) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := domain.UserSearchForm{}
		decoder := json.NewDecoder(req.Body)

		err := decoder.Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users, err := app.service.SearchUser(req.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// GetUser godoc
// @Summary     верификация входных данных.
// @Description верификация входных данных, проверка что приходит(телефон или email), проверка если такие данные в бд,
// @Description и если есть такой телефон то, отправка смс. А если нет телефона или email, то возвращает ошибку.
// @Description а если есть такой email, то возвращает данные такого пользователя.
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.UserForm  true  "JSON"
// @Success      200  {object} domain.UserForm
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      409  {object}  Error
// @Failure      500  {object}  Error
// @Router      /{id} [GET].
func GetUser(app *App) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		user, err := app.service.GetUser(req.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// CreateUser godoc
// @Summary     создание пользователя
// @Description создание пользователя
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.UserForm  true  "JSON"
// @Success      200  {object} domain.UserForm
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      409  {object}  Error
// @Failure      500  {object}  Error
// @Router      /create [POST].
func CreateUser(app *App) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user domain.User

		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userCreated, err := app.service.CreateUser(req.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userCreated)
	}
}

// UpdateUser godoc
// @Summary     верификация входных данных.
// @Description верификация входных данных, проверка что приходит(телефон или email), проверка если такие данные в бд,
// @Description и если есть такой телефон то, отправка смс. А если нет телефона или email, то возвращает ошибку.
// @Description а если есть такой email, то возвращает данные такого пользователя.
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.UserForm  true  "JSON"
// @Success      200  {object} domain.UserForm
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      409  {object}  Error
// @Failure      500  {object}  Error
// @Router      /{id} [PUT].
func UpdateUser(app *App) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var user domain.UserForm
		err = json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updateUser, err := app.service.UpdateUser(req.Context(), id, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updateUser)
	}
}

// DeleteUser godoc
// @Summary     Удаление пользователя
// @Description удаление пользователя по id
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.UserForm  true  "JSON"
// @Success      200  {object} domain.UserForm
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      409  {object}  Error
// @Failure      500  {object}  Error
// @Router      /{id} [DELETE].
func DeleteUser(app *App) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var userForm domain.UserForm
		if err := json.NewDecoder(req.Body).Decode(&userForm); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = app.service.DeleteUser(req.Context(), id, userForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

}
