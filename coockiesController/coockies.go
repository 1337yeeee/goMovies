package coockiesController

import (
	"net/http"
	"strconv"
	"log"

	"movies_crud/structs"
	h "movies_crud/helper"
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
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	cookie, err := r.Cookie("user_id")
	if err == nil {
		if cookie.Value != "" {
			id, err := strconv.Atoi(cookie.Value)
			if err != nil {
				logger.Printf("coockiesController.GetUserCookie(); strconv.Atoi()| %v\n", err)
			}

			user, err := structs.GetUser(id)

			if err != nil {
				logger.Printf("coockiesController.GetUserCookie(); structs.GetUser(id=%v)| %v\n", id, err)
				DelUserCookie(w)
			}

			return &user
		} else {
			return nil
		}
	} else {
		logger.Printf("coockiesController.GetUserCookie(); r.Cookie(`user_id`)| %v\n", err)
		return nil
	}
}

func GetUserCookieIDonly(r *http.Request) int {
	logger, logFile, err := h.CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseLogger(logFile)

	cookie, err := r.Cookie("user_id")
	if err == nil {
		if cookie.Value != "" {
			id, err := strconv.Atoi(cookie.Value)
			if err != nil {
				logger.Printf("coockiesController.GetUserCookieIDonly(); strconv.Atoi()| %v\n", err)
			}

			return int(id)
		} else {
			return 0
		}
	} else {
		logger.Printf("coockiesController.GetUserCookieIDonly(); r.Cookie(`user_id`)| %v\n", err)
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