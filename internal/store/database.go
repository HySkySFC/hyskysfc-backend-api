package store

import (
	"fmt"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=hyskysfc password=hyskysfc dbname=hyskysfc port=5432 sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to database...")

	return db, nil
}
