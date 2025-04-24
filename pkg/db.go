package pkg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	// create database connection string
	dbEnv := []any{}
	dbEnv = append(dbEnv, os.Getenv("DBUSER"))
	dbEnv = append(dbEnv, os.Getenv("DBPASS"))
	dbEnv = append(dbEnv, os.Getenv("DBHOST"))
	dbEnv = append(dbEnv, os.Getenv("DBPORT"))
	dbEnv = append(dbEnv, os.Getenv("DBNAME"))
	// setup database connection
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv...)
	var err error
	DB, err = pgxpool.New(context.Background(), dbString)
	if err != nil {
		return err
	}
	// test oing db
	err = DB.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	log.Println("Connected to PostgreSQL")
	return nil
}
