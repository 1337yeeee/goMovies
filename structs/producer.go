package structs

import (
	"database/sql"
	"movies_crud/data"
)

type Producer struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Img sql.NullString `sql:"img"`
	Description sql.NullString `sql:"description"`
}

func GetProducer(id int) (Producer, error) {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, img, description FROM producers WHERE id = ?", id)

	producer := Producer{}
	err := row.Scan(&producer.ID, &producer.Name, &producer.Img, &producer.Description)

	return producer, err
}