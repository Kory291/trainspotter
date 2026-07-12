package database

import (
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Train struct{
	Tz int				`json:"tz"`
	Baureihe string		`json:"baureihe"`
	Name *string		`json:"name"`
}

func GetTrainsFromDB() (trains []Train, err error) {
	
	rows, err := queryDB("SELECT tz, baureihe, name FROM trains;")
	if err != nil {
		fmt.Printf("There was an error when querying trains: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	trains, err = pgx.CollectRows(rows, func(row pgx.CollectableRow) (Train, error) {
		var t Train
		err := row.Scan(&t.Tz, &t.Baureihe, &t.Name)
		return t, err
	})

	fmt.Printf("Found trains %v\n", trains)
	if err != nil {
		fmt.Printf("There was an error when parsing the DB response to get trains %v\n", err)
		return nil, err
	}
	return trains, nil
}

func AddTrainToDB(train Train) (err error) {
	var name string
	if train.Name == nil {
		name = ""
	} else {
		name = *train.Name
	}
	formattedQuery := fmt.Sprintf("INSERT INTO trains (tz, baureihe, name) VALUES (%d, '%s', '%s')", train.Tz, train.Baureihe, name)
	fmt.Printf("Applying ... %s", formattedQuery)
	rows, err := queryDB(formattedQuery)
	if err != nil {
		rows.Close()
		fmt.Printf("There was an error when inserting into the DB %v\n", err)
		return err
	}
	defer rows.Close()

	output, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var out string
		err := row.Scan(out)
		return out, err
	})
	fmt.Printf("Output of adding stuff to DB ... %s\n", output)
	if err != nil {
		fmt.Printf("There was a second error when inserting train data into the DB %v\n", err)
		return err
	}
	return nil
}