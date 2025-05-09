package pkg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type MyDB struct {
	DBString string
}

func InitDB() *MyDB {
	// create database connection string
	dbEnv := []any{}
	dbEnv = append(dbEnv, os.Getenv("DBUSER"))
	dbEnv = append(dbEnv, os.Getenv("DBPASS"))
	dbEnv = append(dbEnv, os.Getenv("DBHOST"))
	dbEnv = append(dbEnv, os.Getenv("DBPORT"))
	dbEnv = append(dbEnv, os.Getenv("DBNAME"))
	// setup database connection
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv...)
	return &MyDB{
		DBString: dbString,
	}
}

func (m *MyDB) Connect() (*pgxpool.Pool, error) {
	var err error
	DB, err := pgxpool.New(context.Background(), m.DBString)
	if err != nil {
		return nil, err
	}

	// test ping db
	err = DB.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}
	log.Println("Connected to PostgreSQL")

	return DB, nil
}

func (m *MyDB) Migrate() error {
	// Migrasi
	dbConfig, err := pgxpool.ParseConfig(m.DBString)
	if err != nil {
		return err
	}
	stdlib.RegisterConnConfig(dbConfig.ConnConfig)
	sqlDB := stdlib.OpenDB(*dbConfig.ConnConfig)

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		return err
	}

	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
