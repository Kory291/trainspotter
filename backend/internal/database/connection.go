package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

func connect() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}

func queryDB(query string) (pgx.Rows, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	
	conn, err := connect()
	if err != nil {
		fmt.Println("Couldnt connect to DB")
		return nil, err
	}

	rows, err := conn.Query(ctx, query)
	if err != nil {
		fmt.Printf("had a problem querying the DB %v\n", err)
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}