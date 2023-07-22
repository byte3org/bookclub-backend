package userhandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/byte3/bookclub/backend/helpers"
	"github.com/byte3/bookclub/backend/helpers/jwt"
	"github.com/byte3/bookclub/backend/internal/database"
	"github.com/byte3/bookclub/backend/internal/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := database.SelectAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password, err = helpers.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "password cannot be stored", http.StatusInternalServerError)
		return
	}

	database.InsertUser(&user)
	msg := map[string]interface{}{
		"id":   user.Id,
		"name": user.Username,
	}

	j, err := json.Marshall(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(j))
	return
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err = json.Unmarshal(body, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashed_password, err := helpers.HashPassword(u.Password)
	if err != nil {
		http.Error(w, "password cannot be stored", http.StatusInternalServerError)
		return
	}

	// check if user is in the database
	// search for username in the db
	user, err := database.SelectUserbyName(u.Username)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	if user.Username != u.Username || user.Password != hashed_password {
		http.Error(w, "authentication failed", http.StatusBadRequest)
		return
	}

	token_str := jwt.GenerateToken(user.Username, user.Email, user.Password)
	msg := map[string]interface{}{
		"status": "user authenticated",
		"token":  token_str,
	}

	w.WriteHeader(200)
	w.Write([]byte(msg))
	return
}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user, err := database.GetUserDetails(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
