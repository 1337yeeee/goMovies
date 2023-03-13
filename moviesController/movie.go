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
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	movie_idSTR := r.URL.Query().Get("id")
	if movie_idSTR == "" {
		response := structs.Response{}
		response.User = cookie.GetUserCookie(w, r)
		response.Title = "NOT FOUND"
		h.Templating(w, "notfound", "base", response)
		return
	}

	movie_id, err := strconv.Atoi(movie_idSTR)
	if err != nil {
		logger.Printf("moviesController.MovieIndexHandler(); strconv.Atoi()| %v\n", err)
	}
	movie, err := structs.GetMovie(movie_id)
	if err != nil {
		logger.Printf("moviesController.MovieIndexHandler(); structs.GetMovie()| %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	movie.CountRating()

	resp := Response{}
	resp.Movie = &movie
	resp.User = cookie.GetUserCookie(w, r)
	if resp.User != nil {
		resp.Movie.SetUserRating(resp.User.ID)
		resp.IsMovieWatched = structs.IsMovieWatched(resp.User.ID, movie_id)
	} else {
		resp.Movie.SetUserRating(0)
	}
	
	h.Templating(w, "movie", "base", resp)
}

func MoviesIndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	resp := Response{}
	resp.User = cookie.GetUserCookie(w, r)

	var movies []Movie

	if resp.User != nil {
		movies, err = structs.GetMovies(resp.User.ID)
	} else {
		movies, err = structs.GetMovies(0)
	}
	if err != nil {
		logger.Printf("moviesController.MoviesIndexHandler(); structs.GetMovies(isUser=%b)| %v\n", resp.User != nil, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resp.Movies = movies

	h.Templating(w, "movies", "base", resp)
}

func Watched(w http.ResponseWriter, r *http.Request) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	userID := cookie.GetUserCookieIDonly(r)

	movieIDStr := r.URL.Query().Get("movie_id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		logger.Printf("moviesController.Watched(); strconv.Atoi()| %v\n", err)
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	watched, err := structs.DealWithWatched(userID, movieID)
	if err != nil {
		logger.Printf("moviesController.Watched(); structs.DealWithWatched()| %v\n", err)
		http.Error(w, "Failed to handle watched status", http.StatusInternalServerError)
		return
	}

	response := struct {
		Watched bool `json:"watched"`
	}{
		Watched: watched,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Printf("moviesController.Watched(); json.NewEncoder(w)| %v\n", err)
	}
}

func Rated(w http.ResponseWriter, r *http.Request) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)
	
	err = r.ParseForm()
	if err != nil {
		logger.Printf("moviesController.Rated(); r.ParseForm()| %v\n", err)
	}

	movie_id, err := strconv.Atoi(r.Form.Get("movie_id"))
	if err != nil {
		logger.Printf("moviesController.Rated(); strconv.Atoi(movie_id)| %v\n", err)
		return
	}
	rating, err := strconv.Atoi(r.Form.Get("rating"))
	if err != nil {
		logger.Printf("moviesController.Rated(); strconv.Atoi(rating)| %v\n", err)
		return 
	}

	userID := cookie.GetUserCookieIDonly(r)

	structs.SetMovieRating(userID, movie_id, rating)

	if err != nil {
		logger.Printf("moviesController.Rated(); structs.SetMovieRating()| %v\n", err)
		return 
	}

	response := struct {
		Rating int `json:"rating"`
	}{
		Rating: rating,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Printf("moviesController.Rated(); json.NewEncoder(w)| %v\n", err)
	}
}
