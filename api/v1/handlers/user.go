package userhandler

import (
	"net/http"

	"github.com/byte3/bookclub/backend/helpers"
	"github.com/byte3/bookclub/backend/internal/database"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	hashed_password, err := helpers.HashPassword(password)
	if err != nil {
		http.Error(w, "password cannot be stored", http.StatusInternalServerError)
		return
	}

	database.InsertUser(username, email, hashed_password)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("user created"))
	return
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	hashed_password, err := helpers.HashPassword(password)
	if err != nil {
		http.Error(w, "password cannot be stored", http.StatusInternalServerError)
		return
	}

	// check if user is in the database
	// search for username in the db
	user, err := database.SelectUserbyName(username)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	if user.Username != username || user.Password != hashed_password {
		http.Error(w, "authentication failed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("user authenticated"))

	// write auth token
	w.Write([]byte())

	return
}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
