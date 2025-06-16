package main

import (
	"log"
	"net/http"

	"gitlab.com/kovarniykrab/servermain/server/api/routers"
	"gitlab.com/kovarniykrab/servermain/server/database"
)

func main() {
	// Инициализация подключения к базе данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	router := routers.SetupRoutes(db)

	// Запуск сервера
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
