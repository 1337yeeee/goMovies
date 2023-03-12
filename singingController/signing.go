package singingController

import (
	"database/sql"
	"net/http"
	"log"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"movies_crud/structs"
	cookie "movies_crud/coockiesController"
	h "movies_crud/helper"
)

type User = structs.User

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.Templating(w, "signup", "sign_layout")
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Printf("singingController.SignUpHandler(); r.ParseForm()| %v\n", err)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("password")), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("singingController.SignUpHandler(); bcrypt.GenerateFromPassword()| %v\n", err)
		}

		user := User{
			Name: sql.NullString{String: r.Form.Get("name"), Valid: true},
			Email: sql.NullString{String: r.Form.Get("email"), Valid: true},
			Password: hashedPassword,
		}

		res := user.Add()
		if res {
			cookie.SetUserCookie(w, user.ID)

			http.Redirect(w, r, "./", http.StatusFound)
		} else {
			resp := structs.Response{}
			resp.Message = "user already exists"
			h.Templating(w, "signup", "sign_layout", resp)
		}
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.Templating(w, "signin", "sign_layout")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Printf("singingController.SignInHandler(); r.ParseForm()| %v\n", err)
		}

		userId, err := structs.GetUserIDLogin(r.Form.Get("email"), r.Form.Get("password"))
		if err != nil {
			log.Printf("singingController.SignInHandler(); structs.GetUserIDLogin()| %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if userId == 0 {
			fmt.Println(userId)
		} else {
			user, err := structs.GetUser(userId)
			if err != nil {
				log.Printf("singingController.SignInHandler(); structs.GetUser()| %v\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w = cookie.SetUserCookie(w, user.ID)

			http.Redirect(w, r, "./", http.StatusFound)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie.DelUserCookie(w)
	http.Redirect(w, r, "./", http.StatusFound)
}