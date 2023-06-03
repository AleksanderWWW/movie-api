package main

import (
	"log"
	"movie-api/api/v1/handlers"
	"movie-api/api/v1/router"
	"movie-api/db"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	repo := db.MockRepo{}

	router := router.Initialize()

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/", handlers.GetAllMoviesHandler(&repo))
	})

	log.Fatal(http.ListenAndServe(":5555", router))
}
