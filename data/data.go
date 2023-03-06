package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DbName string

func Init(dbName string) error {
	DbName = dbName
	db, err := sql.Open("sqlite3", DbName)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS producers (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		img TEXT,
		description TEXT
		);

		CREATE TABLE IF NOT EXISTS movies(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		year INTEGER,
		country TEXT,
		description TEXT,
		img TEXT,
		producer_id INTEGER,
		FOREIGN KEY (producer_id) REFERENCES producers(id) ON DELETE CASCADE
		);
	`)
	
	return err
}

func DBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", DbName)
	if err != nil {
		panic(err)
	}
	return db
}
