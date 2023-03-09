package structs

import (
	"movies_crud/data"
)

func DealWithRating(user_id int, movie_id int, rating int) error {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(*) FROM user_movie_ratings WHERE user_id = ? AND movie_id = ?", user_id, movie_id)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		_, err := db.Exec(`UPDATE user_movie_ratings SET rating = ?
			WHERE user_id = ? AND movie_id = ?`, rating, user_id, movie_id)
		return err
	} else {
		_, err := db.Exec(`INSERT INTO user_movie_ratings (user_id, movie_id, rating)
							VALUES (?, ?, ?)`, user_id, movie_id, rating)
		return err
	}
}

func DealWithWatched(user_id int,  movie_id int) (bool, error) {
	db := data.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(*) FROM user_movie_ratings WHERE user_id = ? AND movie_id = ?", user_id, movie_id)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count != 0 {
		_, err = db.Exec("UPDATE user_movie_ratings SET watched = NOT watched WHERE user_id = ? AND movie_id = ?", user_id, movie_id)
		if err != nil {
			return false, err
		}

		var watched bool
		err = db.QueryRow("SELECT watched FROM user_movie_ratings WHERE user_id = ? AND movie_id = ?", user_id, movie_id).Scan(&watched)
		if err != nil {
			return false, err
		}
		return watched, nil
	} else {
		_, err = db.Exec("INSERT INTO user_movie_ratings (user_id, movie_id, watched) VALUES (?, ?, ?)", user_id, movie_id, true)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func IsMovieWatched(user_id int, movie_id int) bool {
	db := data.DBConnection()
	defer db.Close()

	var watched bool
	err := db.QueryRow("SELECT watched FROM user_movie_ratings WHERE user_id = ? AND movie_id = ?", user_id, movie_id).Scan(&watched)
	if err != nil {
		return false
	}

	return watched
}