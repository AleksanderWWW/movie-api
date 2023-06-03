package router

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

)

func Initialize() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.RedirectSlashes, 
		middleware.Recoverer, //middleware to recover from panics
		middleware.Heartbeat("/health"), //for heartbeat process such as Kubernetes liveprobeness
		
	)

	//Sets context for all requests
	router.Use(middleware.Timeout(30 * time.Second))

	return router
}
