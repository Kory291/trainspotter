package api

import (
	"net/http"
)

func GetTrains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`Will return a list of trains in JSON format`))
}

func PostTrain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`Will add a new train and return the created train in JSON format`))
}