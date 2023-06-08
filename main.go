package main

import (
	"log"
	"movie-api/api/v1/handlers"
	"movie-api/api/v1/router"
	"movie-api/db"
	"net/http"
)

func main() {
	repo, err := db.NewSqliteRepo()

	if err != nil {
		log.Fatal(err)
	}

	router := router.Initialize()

	router.Route("/api/v1", handlers.Routes(repo))

	log.Fatal(http.ListenAndServe(":5555", router))
}
