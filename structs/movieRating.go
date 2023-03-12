package structs

import (
	"fmt"
	"log"

	"movies_crud/data"
)

func GetMovieRating(movie_id int) (string, error) {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT AVG(rating) FROM user_movie_ratings WHERE movie_id = ? AND rating > 0", movie_id)

	var rating float64
	err := row.Scan(&rating)
	if err != nil {
		log.Printf("structs.GetMovieRating(); row.Scan()| %v\n", err)
	}

	return fmt.Sprintf("%.1f", rating), err
}

func SetMovieRating(user_id int, movie_id int, rating int) error {
	db := data.DBConnection()
	defer db.Close()

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM user_movie_ratings WHERE user_id = ? AND movie_id = ?", user_id, movie_id)
	err := row.Scan(&count)
	if err != nil {
		log.Printf("structs.SetMovieRating(); row.Scan()| %v\n", err)
		return err
	}

	if count == 0 {
		_, err := db.Exec(`INSERT INTO user_movie_ratings (user_id, movie_id, rating) VALUES (?,?,?)`, user_id, movie_id, rating)
		if err != nil {
			log.Printf("structs.SetMovieRating(); db.Exec(`INSERT`)| %v\n", err)
		}
		return err
	} else {
		_, err := db.Exec(`UPDATE user_movie_ratings SET rating = ? WHERE user_id = ? AND movie_id = ?`, rating, user_id, movie_id)
		if err != nil {
			log.Printf("structs.SetMovieRating(); db.Exec(`UPDATE`)| %v\n", err)
		}
		return err
	}
}