package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trainspotter-backend/internal/database"
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
	var sighting database.Sighting
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("There went something wrong: %v", err)
		http.Error(w, "Body was malformed", 400)
	}
	json.Unmarshal(requestBody, &sighting)
	database.AddSightingToDB(sighting)
}