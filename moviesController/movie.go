package moviesController

import (
	"net/http"
	"strconv"

	"movies_crud/structs"
	h "movies_crud/helper"
	cookie "movies_crud/coockiesController"
)

type Movie = structs.Movie
type Response = structs.Response

func MovieIndexHandler(w http.ResponseWriter, r *http.Request) {
	id_ := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(id_)
	movie, err := structs.GetMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := Response{}
	resp.Movie = &movie
	resp.User = cookie.GetUserCookie(w, r)
	
	h.Templating(w, "movie", "base", resp)
}

func MoviesIndexHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{}
	resp.User = cookie.GetUserCookie(w, r)

	movies, err := structs.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resp.Movies = movies

	h.Templating(w, "movies", "base", resp)
}