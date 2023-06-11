package handlers

import (
	"encoding/json"
	"fmt"
	"movie-api/db"
	"movie-api/internal/movie"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Connect new enpoints here
func Routes(repo db.Repo) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/movies", GetAllMoviesHandler(repo))
		r.Get("/movies/{movieID}", GetMovieByIDHandler(repo))
		r.Post("/movies/add", CreateMovieHandler(repo))
		r.Put("movies/update", UpdateMovieHandler(repo))
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
			json.NewEncoder(w).Encode(map[string]string{"msg": "Error parsing id"})
			return
		}

		movie := repo.GetMovieByID(int_id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movie)
	}
}

func CreateMovieHandler(repo db.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		var singleMovie movie.Movie
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&singleMovie)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}

		err = repo.CreateMovie(singleMovie)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"msg": "movie '" + singleMovie.Title + "' created"})
	}
}

func UpdateMovieHandler(repo db.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		var singleMovie movie.Movie
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&singleMovie)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}

		err = repo.UpdateMovie(singleMovie)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": fmt.Sprintf("Movie with ID '%d' updated", singleMovie.ID),
		})
	}
}
