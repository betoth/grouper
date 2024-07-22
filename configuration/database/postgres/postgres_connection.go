package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(ctx context.Context) (*sql.DB, error) {

	ConnStr := "host=localhost port=5432 user=groupe password=groupe dbname=groupe sslmode=disable"
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
