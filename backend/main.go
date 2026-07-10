package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

	"trainspotter-backend/internal/api"
)

var routes = func(r chi.Router) {
	r.Get("/trains", api.GetTrains)
	r.Post("/trains", api.PostTrain)
	r.Get("/sightings", api.GetSightings)
	r.Post("/sightings", api.PostSighting)
}


func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
	routes(r)
    http.ListenAndServe(":8080", r)
}