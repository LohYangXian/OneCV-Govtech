package db

import (
	"database/sql"
	"fmt"
)

func Connect(url string) {
	// connect to database

	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Printf("Error connecting to database, %v", err)
		return
	}

	// Defer the close till after the main function has finished
	defer db.Close()
}
