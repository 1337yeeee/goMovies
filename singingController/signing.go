package singingController

import (
	"database/sql"
	"net/http"
	"regexp"
	"log"

	"golang.org/x/crypto/bcrypt"
	"movies_crud/structs"
	cookie "movies_crud/coockiesController"
	h "movies_crud/helper"
)

type User = structs.User

func ValidateEmail(email string) bool {
	emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegexp.MatchString(email) {
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	passwordRegexp := regexp.MustCompile(`^[a-zA-Z0-9]*[a-zA-Z][0-9]+[a-zA-Z0-9]*$`)
	if !passwordRegexp.MatchString(password) {
		return false
	}
	return true
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	if r.Method == "GET" {
		h.Templating(w, "signup", "sign_layout")
	} else {
		err := r.ParseForm()
		if err != nil {
			logger.Printf("singingController.SignUpHandler(); r.ParseForm()| %v\n", err)
		}

		isValidPass := ValidatePassword(r.Form.Get("password"))
		if ! isValidPass {
			resp := structs.Response{}
			resp.Message = "password should be at least 8 characters long and contain at least one letter and one digit"
			h.Templating(w, "signup", "sign_layout", resp)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("password")), bcrypt.DefaultCost)
		if err != nil {
			logger.Printf("singingController.SignUpHandler(); bcrypt.GenerateFromPassword()| %v\n", err)
		}

		isValidEmail := ValidateEmail(r.Form.Get("email"))
		if ! isValidEmail {
			resp := structs.Response{}
			resp.Message = "invalid email"
			h.Templating(w, "signup", "sign_layout", resp)
			return
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
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	if r.Method == "GET" {
		h.Templating(w, "signin", "sign_layout")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			logger.Printf("singingController.SignInHandler(); r.ParseForm()| %v\n", err)
		}

		isValidEmail := ValidateEmail(r.Form.Get("email"))
		if ! isValidEmail {
			resp := structs.Response{}
			resp.Message = "invalid email"
			h.Templating(w, "signin", "sign_layout", resp)
			return
		}
		isValidPass := ValidatePassword(r.Form.Get("password"))
		if ! isValidPass {
			resp := structs.Response{}
			resp.Message = "password should be at least 8 characters long and contain at least one letter and one digit"
			h.Templating(w, "signin", "sign_layout", resp)
			return
		}

		userId, err := structs.GetUserIDLogin(r.Form.Get("email"), r.Form.Get("password"))
		if err != nil {
			logger.Printf("singingController.SignInHandler(); structs.GetUserIDLogin()| %v\n", err)

			resp := structs.Response{}
			resp.Message = "inccorrect email or password"
			h.Templating(w, "signin", "sign_layout", resp)
			return
		}
		if userId != 0 {
			user, err := structs.GetUser(userId)
			if err != nil {
				logger.Printf("singingController.SignInHandler(); structs.GetUser()| %v\n", err)
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