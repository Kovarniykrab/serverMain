package router

import (
	"github.com/gorilla/mux"

	"gitlab.com/kovarniykrab/servermain/server/api/handlers"
	"gitlab.com/kovarniykrab/servermain/server/config"
	"gitlab.com/kovarniykrab/servermain/server/database"
)

func SetupRoutes(db database.Database) *mux.Router {
	router := mux.NewRouter()
	userHandler := &handlers.UserHandler{}

	// Инициализация обработчиков

	// Маршруты для работы с пользователями
	router.HandleFunc("/users", userHandler.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	return router
}

type App struct {
	Config     config.Config
	I18n       i18n.Translator
	Dictionary *dictionary.Dictionary
	Service    *service.App
}
