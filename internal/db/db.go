package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Driver para PostgreSQL
)

func Connect() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database: %w", err)
	}

	fmt.Println("Connected to PostgreSQL!")
	return db, nil
}
