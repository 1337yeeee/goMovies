package main

import (
	"net/http"
	"fmt"

	"movies_crud/data"
	"movies_crud/structs"
	singing "movies_crud/singingController"
	cookie "movies_crud/coockiesController"
	h "movies_crud/helper"
)

const PORT = "8000"

type User = structs.User

func indexHandler(w http.ResponseWriter, r *http.Request) {
	user := cookie.GetUserCookie(r)
	if user != nil {
		h.Templating(w, "index", "base", *user)
	} else {
		h.Templating(w, "index", "base")
	}
}

func main() {
	data.Init("test.db")
	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/signup", singing.SignUpHandler)
	mux.HandleFunc("/login", singing.SignInHandler)
	mux.HandleFunc("/logout", singing.LogoutHandler)
	fmt.Printf("Listening on :%v\n", PORT)
	http.ListenAndServe(":"+PORT, mux)
}
