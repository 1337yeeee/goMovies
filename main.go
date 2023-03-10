package main

import (
	"net/http"
	"fmt"
	"log"
	"os"

	"movies_crud/data"
	"movies_crud/structs"
	singing "movies_crud/singingController"
	cookie "movies_crud/coockiesController"
	movie "movies_crud/moviesController"
	h "movies_crud/helper"

	"github.com/joho/godotenv"
)

type User = structs.User
type Response = structs.Response

func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{}
	resp.User = cookie.GetUserCookie(w, r)
	h.Templating(w, "index", "base", resp)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	LogFileName := os.Getenv("LOGFILE")
	if LogFileName == "" {
		LogFileName = "tmp.log"
	}
	h.SetLogFileName(LogFileName)
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	dbName := os.Getenv("DB_FILE_NAME")
	if dbName == "" {
		dbName = "tmp.db"
	}

	err = data.Init(dbName)
	if err != nil {
		logger.Printf("main; data.Init()| %v\n", err)
	}
	h.CloseLogger(logFile)


	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/signup", singing.SignUpHandler)
	mux.HandleFunc("/login", singing.SignInHandler)
	mux.HandleFunc("/logout", singing.LogoutHandler)
	mux.HandleFunc("/movie", movie.MovieIndexHandler)
	mux.HandleFunc("/director", movie.DirectorIndexHandler)
	mux.HandleFunc("/movies", movie.MoviesIndexHandler)
	mux.HandleFunc("/rate", movie.Rated)
	mux.HandleFunc("/watched", movie.Watched)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Listening on :%v\n", port)
	http.ListenAndServe(":"+port, mux)
}
