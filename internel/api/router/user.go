package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/kovarniykrab/servermain/internel/api/handlers"
)

func (app *App) SetupUserRoutes(router *mux.Router) {

	router.HandleFunc("", handlers.SearchUser(&app.App)).Methods(http.MethodPost)
	router.HandleFunc("/{id}", handlers.GetUser(&app.App)).Methods("GET")
	router.HandleFunc("/create", handlers.CreateUser(&app.App)).Methods("POST")
	router.HandleFunc("/{id}", handlers.UpdateUser(&app.App)).Methods("PUT")
	//router.HandleFunc("/users/{id}", handlers.DeleteUser(&app.App)).Methods("DELETE")

}
