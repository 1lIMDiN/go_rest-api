package api

import (
	"encoding/json"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/pkg/db"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method format, want: GET", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(db.Tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}