package requestshandler

import (
	"encoding/json"
	"net/http"

	"github.com/byte3/bookclub/backend/internal/database"
)

func GetAllRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

}

func GetAllRequestsCount(w http.ResponseWriter, r *http.Request) {

}

func GetAllPendingRequests(w http.ResponseWriter, r *http.Request) {

}

func GetAllAcceptedRequests(w http.ResponseWriter, r *http.Request) {

}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
