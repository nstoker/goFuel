package database

import (
	"database/sql"
	"fmt"
	"os"

	m "github.com/nstoker/gofuel/internal/pkg/migrations"
)

var (
	dsn string
	db  *sql.DB
)

// InitialiseDatabase initialises the database.
func InitialiseDatabase() error {
	var err error

	dsn = os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("database url environment variable missing")
	}

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return m.MigrateUp(dsn)
}
