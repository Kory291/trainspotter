package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"trainspotter-backend/internal/database"
)


func GetTrains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	trains, err := database.GetTrainsFromDB()
	if err != nil {
		fmt.Printf("There was an error when getting the trains from the DB %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonMessage, err := json.Marshal(trains)
	if err != nil {
		fmt.Printf("There was an error when enconding the DB response in json ... %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonMessage)
}

func PostTrain(w http.ResponseWriter, r *http.Request) {
	var train database.Train

	tz := r.FormValue("tz")
	if tz == "" {
		fmt.Printf("There was no tz detected\n")
		http.Error(w, "Malformed Data was provided", 400)
		return
	}
	tzInt, err := strconv.Atoi(tz)
	if err != nil {
		fmt.Printf("There was an error ... %v", err)
		http.Error(w, "Malformed Data was provided", 400)
		return
	}
	train.Tz = tzInt

	train.Baureihe = r.FormValue("baureihe")
	name := r.FormValue("name")
	train.Name = &name

	err = database.AddTrainToDB(train)
	if err != nil {
		fmt.Printf("There was something wrong when running the Post to DB function %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}