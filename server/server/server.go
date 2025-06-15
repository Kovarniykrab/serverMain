package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Разные методы для одного пути
	r.HandleFunc("/item", handleItem).Methods("GET")
	r.HandleFunc("/item", handleItem).Methods("POST")
	r.HandleFunc("/item", handleItem).Methods("DELETE")

	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", r)
}

func handleItem(res http.ResponseWriter, req http.Request) {
	switch req.Method {
	case "POST":
		var request struct {
			Name string `json: "name"`
		}

		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			http.Error(res, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if request.Name == "" {
			http.Error(res, "Name is required", http.StatusBadRequest)
			return
		}

		id, err := database.CreateItem(request.Name)
		if err != nil {
			http.Error(res, "Database Error", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(map[string]interface{}{
			"id":   id,
			"name": request.Name,
		})

	case "GET":
		idStr := req.URL.Query().Get("id")

		// Если есть ID - возвращаем одну запись
		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(res, "Invalid ID format", http.StatusBadRequest)
				return
			}

			item, err := database.GetItem(id)
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					http.Error(res, "Item not found", http.StatusNotFound)
				} else {
					http.Error(res, "Database error", http.StatusInternalServerError)
				}
				return
			}

			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(item)
			return
		}

		// Если нет ID - возвращаем все записи
		items, err := database.GetAllItems()
		if err != nil {
			http.Error(res, "Database error", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(items)

	case "DELETE":

		idStr := req.URL.Query().Get("id")
		if idStr == "" {
			http.Error(res, "ID parameter is required", http.StatusBadRequest)
			return
		}

		// Конвертируем ID в число
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(res, "Invalid ID format", http.StatusBadRequest)
			return
		}

		// Удаляем запись
		err = database.DeleteItem(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				http.Error(res, err.Error(), http.StatusNotFound)
			} else {
				http.Error(res, "Database error", http.StatusInternalServerError)
			}
			return
		}

		// Успешный ответ
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": fmt.Sprintf("Item with ID %d deleted", id),
		})
	}
}
