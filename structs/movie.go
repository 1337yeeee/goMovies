package structs

import (
	"fmt"
	"log"

	"database/sql"
	"movies_crud/data"
	h "movies_crud/helper"
)

type Movie struct {
	ID int `sql:"id"`
	Name sql.NullString `sql:"name"`
	Year int `sql:"year"`
	Country sql.NullString `sql:"country"`
	Description sql.NullString `sql:"description"`
	Img sql.NullString `sql:"img"`
	Director Director
	Rating string
	UserRating int
}

func GetMovie(id int) (Movie, error) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, year, description, img, country, director_id FROM movies WHERE id = ?", id)

	movie := Movie{}
	var director_id int
	err = row.Scan(&movie.ID, &movie.Name, &movie.Year, &movie.Description, &movie.Img, &movie.Country, &director_id)
	if err != nil {
		logger.Printf("structs.GetMovie(); row.Scan()| %v\n", err)
	}
	if director_id != 0 {
		director, _ := GetDirector(director_id)
		movie.Director = director
	}

	return movie, err
}

func GetMovies(user_id int) ([]Movie, error) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)
	
	var movies []Movie

	db := data.DBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, year, description, img FROM movies ORDER BY year DESC")
	if err != nil {
		logger.Printf("structs.GetMovies(); db.Query(`SELECT`)| %v\n", err)
		return movies, err
	}
	defer rows.Close()

	for rows.Next() {
		movie := Movie{}
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Year, &movie.Description, &movie.Img)
		if err != nil {
			logger.Printf("structs.GetMovies(); rows.Scan()| %v\n", err)
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

func GetMoviesOfDirector(director_id int) []Movie {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)
	
	var movies []Movie

	db := data.DBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, year FROM movies WHERE director_id = ? ORDER BY year DESC", director_id)
	if err != nil {
		logger.Printf("structs.GetMovies(); db.Query(`SELECT`)| %v\n", err)
		return movies
	}
	defer rows.Close()

	for rows.Next() {
		movie := Movie{}
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Year)
		if err != nil {
			logger.Printf("structs.GetMovies(); rows.Scan()| %v\n", err)
			return movies
		}
		movies = append(movies, movie)
	}

	return movies
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
