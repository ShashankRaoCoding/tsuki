package initialise 

import (
	"log" 
	"database/sql"
)

// this file initialises databases required by tsuki
const Schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	token TEXT NOT NULL UNIQUE,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
`


func DBSchema( *sql.DB) error {
	var err error 

	// initialise schema for the database
	err = .Exec(Schema)
	if err != nil {
		log.Println(err.Error())
	}

	return err 
}






































