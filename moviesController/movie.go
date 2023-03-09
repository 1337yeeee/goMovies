package moviesController

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"

	"movies_crud/structs"
	h "movies_crud/helper"
	cookie "movies_crud/coockiesController"
)

type Movie = structs.Movie
type Response = structs.Response

func MovieIndexHandler(w http.ResponseWriter, r *http.Request) {
	movie_idSTR := r.URL.Query().Get("id")

	movie_id, _ := strconv.Atoi(movie_idSTR)
	movie, err := structs.GetMovie(movie_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	movie.CountRating()

	resp := Response{}
	resp.Movie = &movie
	resp.User = cookie.GetUserCookie(w, r)
	resp.Movie.SetUserRating(resp.User.ID)
	resp.IsMovieWatched = structs.IsMovieWatched(resp.User.ID, movie_id)
	
	h.Templating(w, "movie", "base", resp)
}

func MoviesIndexHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{}
	resp.User = cookie.GetUserCookie(w, r)

	movies, err := structs.GetMovies(resp.User.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resp.Movies = movies

	h.Templating(w, "movies", "base", resp)
}

func Watched(w http.ResponseWriter, r *http.Request) {
	userID := cookie.GetUserCookieIDonly(r)

	movieIDStr := r.URL.Query().Get("movie_id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	watched, err := structs.DealWithWatched(userID, movieID)
	if err != nil {
		http.Error(w, "Failed to handle watched status", http.StatusInternalServerError)
		return
	}

	response := struct {
		Watched bool `json:"watched"`
	}{
		Watched: watched,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Rated(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	movie_id, err := strconv.Atoi(r.Form.Get("movie_id"))
	if err != nil {
		return
	}
	rating, err := strconv.Atoi(r.Form.Get("rating"))
	if err != nil {
		return 
	}

	userID := cookie.GetUserCookieIDonly(r)

	structs.SetMovieRating(userID, movie_id, rating)

	if err != nil {
		return 
	}

	response := struct {
		Rating int `json:"rating"`
	}{
		Rating: rating,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
