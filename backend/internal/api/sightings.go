package api

import (
	"fmt"
	"net/http"
	"trainspotter-backend/internal/database"
	"encoding/json"
)

func GetSightings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sightings, err := database.GetSightingsFromDB()
	if err != nil {
		fmt.Println("There was something wrong when getting the sightings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(sightings)
	if err != nil {
		fmt.Println("Could not convert data to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func PostSighting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`Will add a new sighting and return the created sighting in JSON format`))
}