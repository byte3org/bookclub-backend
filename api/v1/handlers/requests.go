package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/byte3/bookclub/backend/internal/database"
	"github.com/byte3/bookclub/backend/internal/models"
	"github.com/google/uuid"
)

func GetAllRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := database.SelectAllRequests()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&requests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req models.BookRequestModel
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.InsertRequest(&req)
	msg := map[string]interface{}{
		"id": req.Id,
	}

	j, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(j))
}

func GetAllRequestsCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func GetAllPendingRequests(w http.ResponseWriter, r *http.Request) {
	pending_reqs, err := database.SelectAllPendingRequests()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&pending_reqs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func GetAllAcceptedRequests(w http.ResponseWriter, r *http.Request) {
	accepted_reqs, err := database.SelectAllAcceptedRequests()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&accepted_reqs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func GetRequestDetails(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	request, err := database.SelectRequestById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	database.DeleteRequestById(id)
	w.WriteHeader(http.StatusAccepted)
}
