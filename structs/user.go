package structs

import (
	"log"

	"database/sql"
	"movies_crud/data"
	"golang.org/x/crypto/bcrypt"
	h "movies_crud/helper"
)

type User struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Email sql.NullString `sql:"email"`
	Password []byte `sql:"password"`
}

func (user *User) Add() bool {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	db := data.DBConnection()
	defer db.Close()

	result, err := db.Exec(`INSERT INTO users
		(name, email, password)
		VALUES (?, ?, ?)`, user.Name, user.Email, user.Password)

	if err != nil {
		logger.Printf("structs.(User).Add(): db.Exec(INSERT)| %v\n")
		return false
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Printf("structs.(User).Add(): LastInsertId()| %v\n")
	}

	user.ID = int(id)
	return true
}

func GetUser(id int) (User, error) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	user := User{}
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		logger.Printf("structs.GetUser(): row.Scan()| %v\n")
	}

	return user, err
}

func GetUserIDLogin(email string, password string) (int, error) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	db := data.DBConnection()
	defer db.Close()


	row := db.QueryRow("SELECT id, password FROM users WHERE email = ?", email)

	var id int
	var hashedPassword string
	err = row.Scan(&id, &hashedPassword)
	if err != nil {
		logger.Printf("structs.GetUserIDLogin(): row.Scan()| %v\n")
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logger.Printf("structs.GetUserIDLogin(): bcrypt.CompareHashAndPassword()| %v\n")
		return 0, err
	}

	return id, nil
}