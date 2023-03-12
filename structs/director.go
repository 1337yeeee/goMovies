package structs

import (
	"log"

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
	if err != nil {
		log.Printf("structs.GetDirector(); row.Scan()| %v\n", err)
	}

	return director, err
}