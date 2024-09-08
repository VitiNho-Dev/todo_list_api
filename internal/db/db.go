package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver para PostgreSQL
)

func Connect() (*sql.DB, error) {
	dsn := "user=victor password=@teste123 dbname=taskdb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %w", err)
	}

	return db, nil
}
