package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
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

	query := `
	DROP TABLE IF EXISTS sessions;
 	DROP TABLE IF EXISTS users;
	CREATE TABLE IF NOT  EXISTS users(id INTEGER PRIMARY KEY, username TEXT,email TEXT NOT NULL UNIQUE,hash_password TEXT NOT NULL,mobile_phone TEXT,birth_date TEXT);
	CREATE TABLE IF NOT EXISTS sessions(id INTEGER PRIMARY KEY,user_id INTEGER,token TEXT NOT NULL UNIQUE,expired_date TEXT NOT NULL,FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE);	
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfuly connect to database")
	return db, nil
}
