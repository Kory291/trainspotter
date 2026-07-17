package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"trainspotter-backend/internal/database"
	"time"
	"strconv"
)

type ReturnedSighting struct {
	Train int `json:"train"`
	Place string `json:"place"`
	Date time.Time `json:"date"`
}

func GetSightings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var returnedSightings []ReturnedSighting
	sightings, err := database.GetSightingsFromDB()
	if err != nil {
		fmt.Println("There was something wrong when getting the sightings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, sighting := range sightings {
		returnedSightings = append(returnedSightings, ReturnedSighting{
			Train: sighting.Train,
			Place: sighting.Place,
			Date: sighting.Date,
		})
	}
	jsonResponse, err := json.Marshal(returnedSightings)
	if err != nil {
		fmt.Println("Could not convert data to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func PostSighting(w http.ResponseWriter, r *http.Request) {
	var sighting database.Sighting

	tz := r.FormValue("tz")
	if tz == "" {
		fmt.Printf("Ther was no tz detected\n")
		http.Error(w, "Malformed data was provided", 400)
		return
	}
	tzInt, err := strconv.Atoi(tz)
	if err != nil {
		fmt.Printf("Couldnt convert string %s to integer", tz)
		http.Error(w, "Malformed data was provided", 400)
		return
	}
	sighting.Train = tzInt
	date := r.FormValue("date")
	place := r.FormValue("place")

	parsedDate, err := time.Parse("02.01.2006", date) 
	sighting.Place = place
	sighting.Date = parsedDate


	err = database.AddSightingToDB(sighting)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
