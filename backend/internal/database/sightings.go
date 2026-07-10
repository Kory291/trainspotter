package database

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	
	conn, err := connect()
	if err != nil {
		fmt.Println("Couldnt connect to DB")
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, place, date, train FROM sightings;`)
	if err != nil {
		fmt.Printf("had a problem querying the DB %v\n", err)
		return nil, err
	}
	defer rows.Close()

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