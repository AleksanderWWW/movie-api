package handlers

import (
	"strconv"
	"encoding/json"
	"movie-api/db"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Connect new enpoints here
func Routes(repo db.Repo) func(r chi.Router) { 
	return func(r chi.Router) {
		r.Get("/movies", GetAllMoviesHandler(repo))
		r.Get("/movies/{movieID}", GetMovieByIDHandler(repo))
	}
}

// Define new endpoints here
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

func GetMovieByIDHandler(repo db.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "movieID")

		int_id, err := strconv.Atoi(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Error parsing id"})
			return
		}

		movie := repo.GetMovieByID(int_id)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(movie)
	}
}
