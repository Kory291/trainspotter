package database

import (
	"time"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Sighting struct{
	Id int  		`json:"id"`
	Place string	`json:"place"`  
	Date time.Time	`json:"date"`
	Train int		`json:"train"`	// this should be the TZ number of the train
} 

func GetSightingsFromDB() (sightings []Sighting, err error) {
	
	rows, err := queryDB(`SELECT id, place, date, train FROM sightings;`)
	if err != nil {
		fmt.Printf("There was an error querying the db %v\n", err)
		return nil, err
	}
	sightings, err = pgx.CollectRows(rows, func(row pgx.CollectableRow) (Sighting, error) {
		var s Sighting
		err := row.Scan(&s.Id, &s.Place, &s.Date, &s.Train)
		return s, err
	})
	if err != nil {
		fmt.Printf("There was a problem when parsing the DB response %v/n", err)
		return nil, err
	}
	return sightings, nil
}

func AddSightingToDB(sighting Sighting) (err error) {

	formattedQuery := fmt.Sprintf("INSERT INTO sightings (place, date, train) VALUES ('%v', '%d-%d-%d', '%v')", sighting.Place, sighting.Date.Year(), sighting.Date.Month(), sighting.Date.Day(), sighting.Train)
	fmt.Printf("Applying ... %s", formattedQuery)
	_, err = queryDB(formattedQuery)
	if err != nil {
		fmt.Printf("There was an error when inserting into the DB %v\n", err)
		return err
	}
	return nil
}