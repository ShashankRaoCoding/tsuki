package initialise 

import (
	"log" 
	"database/sql"
)

// this file initialises databases required by tsuki
const Schema = `
CREATE TABLE names (
  id   INTEGER PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE transactions (
  id       INTEGER PRIMARY KEY,
  name_id  INTEGER NOT NULL REFERENCES names(id) ON DELETE CASCADE,
  amount   REAL NOT NULL,
  date     TEXT NOT NULL
);

CREATE TABLE emails (
  id       INTEGER PRIMARY KEY,
  name_id  INTEGER NOT NULL REFERENCES names(id) ON DELETE CASCADE,
  subject  TEXT,
  content  TEXT,
  date     TEXT NOT NULL
);

CREATE TABLE contacts (
  id            INTEGER PRIMARY KEY,
  name_id       INTEGER NOT NULL REFERENCES names(id) ON DELETE CASCADE,
  email_address TEXT,
  phone_number  TEXT
);

CREATE INDEX idx_transactions_name ON transactions(name_id, date);
CREATE INDEX idx_emails_name       ON emails(name_id, date);
CREATE INDEX idx_contacts_name     ON contacts(name_id);
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






































