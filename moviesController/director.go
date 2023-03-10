package moviesController

import (
	"log"
	"net/http"
	"strconv"

	"movies_crud/structs"
	h "movies_crud/helper"
	cookie "movies_crud/coockiesController"
)

func DirectorIndexHandler(w http.ResponseWriter, r *http.Request) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)
	
	director_idSTR := r.URL.Query().Get("id")
	if director_idSTR == "" {
		response := structs.Response{}
		response.User = cookie.GetUserCookie(w, r)
		response.Title = "NOT FOUND"
		h.Templating(w, "notfound", "base", response)
		return
	}

	director_id, _ := strconv.Atoi(director_idSTR)
	director, err := structs.GetDirector(director_id)
	if err != nil {
		logger.Printf("moviesController.Rated(); structs.GetDirector()| %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := structs.Response{}
	response.User = cookie.GetUserCookie(w, r)
	response.Director = &director

	h.Templating(w, "director", "base", response)
}
