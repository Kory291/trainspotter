package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"fmt"

	"trainspotter-backend/internal/api"
)

var routes = func(r chi.Router) {
	r.Get("/trains", api.GetTrains)
	r.Post("/trains", api.PostTrain)
	r.Get("/sightings", api.GetSightings)
	r.Post("/sightings", api.PostSighting)
}

func main() {
	fmt.Println("Starting server ...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	routes(r)
	fmt.Println("registered routes")
	http.ListenAndServe("0.0.0.0:8080", r)
}
