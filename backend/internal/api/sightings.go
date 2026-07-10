package api

import (
	"net/http"
)

func GetSightings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`Will return a list of sightings in JSON format`))
}

func PostSighting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`Will add a new sighting and return the created sighting in JSON format`))
}