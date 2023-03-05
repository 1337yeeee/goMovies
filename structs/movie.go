package structs

import "database/sql"

type Movie struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Year int `sql:"year"`
	Description sql.NullString `sql:"description"`
	Img sql.NullString `sql:"img"`
	Producer Producer `sql:"producer"`
}