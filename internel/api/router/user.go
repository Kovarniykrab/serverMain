package router

import (
	"github.com/gorilla/mux"
	"gitlab.com/kovarniykrab/servermain/internel/api/handlers"
)

func (app *App) SetupUserRoutes(router *mux.Router) {

	router.HandleFunc("/users", handlers.GetAllUsersHandler(&app.App)).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser(&app.App)).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser(&app.App)).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser(&app.App)).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser(&app.App)).Methods("DELETE")

}
