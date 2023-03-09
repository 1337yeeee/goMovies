package structs

import (
	"database/sql"
	"movies_crud/data"
)

type User struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Email sql.NullString `sql:"email"`
	Password sql.NullString `sql:"password"`
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

	row := db.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", email, password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}