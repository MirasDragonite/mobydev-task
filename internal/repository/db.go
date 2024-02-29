package repository

import (
	"database/sql"
	"fmt"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}

	query := `CREATE TABLE IF NOT  EXISTS users(id INTEGER PRIMARY KEY, username TEXT,email TEXT NOT NULL UNIQUE,hash_password TEXT NOT NULL,mobile_phone TEXT,birth_date TEXT);`

	_, err = db.Exec(query)

	fmt.Println("Successfuly connect to database")
	return db, nil
}
