package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HomeHandler обрабатывает корневой маршрут
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our Go API!")
}

// UserHandler демонстрирует работу с GET и POST
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}
		writeJSON(w, map[string]string{"user_id": id})

	case http.MethodPost:
		var body struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Name == "" {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		writeJSON(w, map[string]string{"created": body.Name})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// Вспомогательная функция для JSON-ответов
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
