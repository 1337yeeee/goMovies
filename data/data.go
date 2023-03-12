package data

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbName string

func Init(dbName string) error {
	DbName = dbName
	db, err := sql.Open("sqlite3", DbName)
	if err != nil {
		log.Printf("data.Init(); sql.Open()| %v\n", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS directors (
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
			director_id INTEGER,
			FOREIGN KEY (director_id) REFERENCES directors(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS user_movie_ratings (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			movie_id INTEGER NOT NULL,
			rating INTEGER DEFAULT 0,
			watched BOOLEAN DEFAULT FALSE,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE
		);

	`)
	
	if err != nil {
		log.Printf("data.Init(); db.Exec(`!ALL!`)| %v\n", err)
	}
	return err
}

func DBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", DbName)
	if err != nil {
		log.Printf("data.DBConnection(); sql.Open()| %v\n", err)
		panic(err)
	}
	return db
}

func ExecuteSQLFile(fileName string) error {
	db := DBConnection()
	defer db.Close()

	// Read SQL file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("data.ExecuteSQLFile(); ioutil.ReadFile()| %v\n", err)
		return err
	}

	// Execute SQL statements
	_, err = db.Exec(string(content))
	if err != nil {
		log.Printf("data.ExecuteSQLFile(); db.Exec(content)| %v\n", err)
		return err
	}

	return nil
}
