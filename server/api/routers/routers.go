package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(repo *services.TodoService) *mux.Router {
	router := mux.NewRouter()
	todoHandler := handlers.NewTodoHandler(repo)

	router.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	return router
}
