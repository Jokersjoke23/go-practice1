package movies

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch movies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
