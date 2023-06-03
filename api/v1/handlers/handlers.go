package handlers

import (
	"encoding/json"
	"movie-api/db"
	"net/http"
)


func GetAllMoviesHandler(repo db.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		allMovies := repo.GetAllMovies()

		if len(allMovies) > 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(allMovies)
	}
}
