package structs

import (
	"fmt"
	"database/sql"
	"movies_crud/data"
)

type Movie struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Year int `sql:"year"`
	Country sql.NullString `sql:"country"`
	Description sql.NullString `sql:"description"`
	Img sql.NullString `sql:"img"`
	Producer Producer
	Rating string
	UserRating int
}

func GetMovie(id int) (Movie, error) {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, year, description, img, country, producer_id FROM movies WHERE id = ?", id)

	movie := Movie{}
	var producer_id int
	err := row.Scan(&movie.ID, &movie.Name, &movie.Year, &movie.Description, &movie.Img, &movie.Country, &producer_id)
	if producer_id != 0 {
		producer, _ := GetProducer(producer_id)
		movie.Producer = producer
	}

	return movie, err
}

func GetMovies(user_id int) ([]Movie, error) {
	var movies []Movie

	db := data.DBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, year, description, img FROM movies ORDER BY year DESC")
	if err != nil {
		return movies, err
	}
	defer rows.Close()

	for rows.Next() {
		movie := Movie{}
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Year, &movie.Description, &movie.Img)
		if err != nil {
			return movies, err
		}
		movie.CountRating()
		if user_id != 0 {
			movie.SetUserRating(user_id)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (movie *Movie) CountRating() {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT AVG(rating) FROM user_movie_ratings WHERE movie_id = ? AND rating > 0", movie.ID)

	var rating float64
	row.Scan(&rating)
	
	movie.Rating = fmt.Sprintf("%.1f", rating)
}

func (movie *Movie) SetUserRating(user_id int) {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT rating FROM user_movie_ratings WHERE movie_id = ? AND user_id = ?", movie.ID, user_id)

	var rating int
	row.Scan(&rating)
	
	movie.UserRating = rating
}