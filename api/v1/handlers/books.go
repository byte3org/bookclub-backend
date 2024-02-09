package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/byte3/bookclub/backend/internal/database"
	"github.com/byte3/bookclub/backend/internal/models"
	"github.com/google/uuid"
)

func GetAllAvailableBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.SelectAllAvailableBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.SelectAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := uuid.Parse(r.Context().Value("user_id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var book models.BookModel
	err = json.Unmarshal(body, &book)
	book.OwnerId = id
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.InsertBook(&book)
	msg := map[string]interface{}{
		"id": book.Id,
	}

	j, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(j))
	return
}

func GetBookDetails(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	book, err := database.SelectBookById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func GetBooksByName(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value("name").(string)

	book, err := database.SelectBooksByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func GetAvailableBooksByName(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value("name").(string)

	book, err := database.SelectBooksByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}
