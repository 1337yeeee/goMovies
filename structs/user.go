package structs

import (
	"database/sql"
	"movies_crud/data"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Email sql.NullString `sql:"email"`
	Password []byte `sql:"password"`
}

func (user *User) Add() bool {
	db := data.DBConnection()
	defer db.Close()

	result, err := db.Exec(`INSERT INTO users
		(name, email, password)
		VALUES (?, ?, ?)`, user.Name, user.Email, user.Password)

	if err != nil {
		return false
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	user.ID = int(id)
	return true
}

func GetUser(id int) (User, error) {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	return user, err
}

func GetUserIDLogin(email string, password string) (int, error) {
	db := data.DBConnection()
	defer db.Close()


	row := db.QueryRow("SELECT id, password FROM users WHERE email = ?", email)

	var id int
	var hashedPassword string
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return 0, err
	}

	return id, nil
}