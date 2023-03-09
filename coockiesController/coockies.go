package coockiesController

import (
	"net/http"
	"strconv"

	"movies_crud/structs"
)

type User = structs.User
var user User

func SetUserCookie(w http.ResponseWriter, userID int) http.ResponseWriter {
	cookie := http.Cookie{
		Name:	"user_id",
		Value:	strconv.Itoa(userID),
		Path:	"/",
		// MaxAge:   3600,
	}

	http.SetCookie(w, &cookie)

	return w
}

func GetUserCookie(w http.ResponseWriter, r *http.Request) *User {
	cookie, err := r.Cookie("user_id")
	if err == nil {
		if cookie.Value != "" {
			id, _ := strconv.Atoi(cookie.Value)
			user, err := structs.GetUser(id)

			if err != nil {
				DelUserCookie(w)
			}

			return &user
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func GetUserCookieIDonly(r *http.Request) int {
	cookie, err := r.Cookie("user_id")
	if err == nil {
		if cookie.Value != "" {
			id, _ := strconv.Atoi(cookie.Value)

			return int(id)
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func DelUserCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:	"user_id",
		Value:	"",
		Path:	"/",
		MaxAge:   -1,
	}

	http.SetCookie(w, &cookie)
}