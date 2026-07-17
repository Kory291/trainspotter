package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool     *pgxpool.Pool
	poolOnce sync.Once
	poolErr  error
)

func getPool() (*pgxpool.Pool, error) {
	poolOnce.Do(func() {
		pool, poolErr = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if poolErr != nil {
			fmt.Printf("Unable to connect to database: %v\n", poolErr)
		}
	})
	return pool, poolErr
}

func queryDB(query string) (pgx.Rows, error) {
	conn, err := getPool()
	if err != nil {
		fmt.Println("Couldnt connect to DB")
		return nil, err
	}

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		fmt.Printf("had a problem querying the DB %v\n", err)
		return nil, err
	}

	return rows, nil
}
