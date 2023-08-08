package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Dbconn() *sql.DB {

	// connStr := "user=postgres password=root dbname=complaints sslmode=disable"

	// Open the database connection
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil
	}
	// Ping the database to ensure a successful connection
	err = db.Ping()
	if err != nil {
		return nil
	}
	fmt.Println("Connected to the database!")
	// Close the database connection when done
	return db
}
