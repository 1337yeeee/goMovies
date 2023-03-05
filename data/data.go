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
		password TEXT NOT NULL);
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
