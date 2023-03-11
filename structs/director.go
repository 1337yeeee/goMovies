package structs

import (
	"database/sql"
	"movies_crud/data"
)

type Director struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Img sql.NullString `sql:"img"`
	Description sql.NullString `sql:"description"`
}

func GetDirector(id int) (Director, error) {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, img, description FROM directors WHERE id = ?", id)

	director := Director{}
	err := row.Scan(&director.ID, &director.Name, &director.Img, &director.Description)

	return director, err
}