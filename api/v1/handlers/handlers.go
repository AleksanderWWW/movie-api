package handlers

import (
	"encoding/json"
	"fmt"
	"movie-api/storage"
	"movie-api/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Connect new enpoints here
func Routes(repo storage.Repo) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/movies", GetAllMoviesHandler(repo))
		r.Get("/movie", GetMovieByIDHandler(repo))
		r.Post("/movie/add", CreateMovieHandler(repo))
		r.Put("/movie/update", UpdateMovieHandler(repo))
		r.Delete("/movie/delete", DeleteMovieHandler(repo))
	}
}

// Define new endpoints here
func GetAllMoviesHandler(repo storage.Repo) http.HandlerFunc {
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

func GetMovieByIDHandler(repo storage.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Query().Get("movieID")

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

func CreateMovieHandler(repo storage.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		var singleMovie models.Movie
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

func UpdateMovieHandler(repo storage.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		var singleMovie models.Movie
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

func DeleteMovieHandler(repo storage.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Query().Get("movieID")

		int_id, err := strconv.Atoi(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"msg": "Error parsing id"})
			return
		}

		err = repo.DeleteMovie(int_id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": fmt.Sprintf("Movie with ID '%d' deleted", int_id),
		})
	}
}
