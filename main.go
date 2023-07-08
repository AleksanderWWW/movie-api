package main

import (
	"flag"
	"fmt"
	"log"
	"movie-api/api/v1/handlers"
	"movie-api/api/v1/router"
	"movie-api/storage"
	"net/http"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5555", "api port to listen on")
	flag.Parse()
	
	repo, err := storage.NewSqliteRepo("movie.db")
	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		log.Fatal(err)
	}

	router := router.Initialize()

	router.Route("/api/v1", handlers.Routes(repo))

	fmt.Println("Listening at port:", *listenAddr)

	log.Fatal(http.ListenAndServe(*listenAddr, router))
}
