package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("There was an error when reading request body to add train ... %v\n", err)
		http.Error(w, "Body was malformed", 400)
		return
	}
	json.Unmarshal(requestBody, &train)
	err = database.AddTrainToDB(train)
	if err != nil {
		fmt.Printf("There was something wrong when running the Post to DB function %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}