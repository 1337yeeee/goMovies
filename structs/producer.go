package structs

import "database/sql"

type Producer struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Img sql.NullString `sql:"img"`
	Description sql.NullString `sql:"description"`
}