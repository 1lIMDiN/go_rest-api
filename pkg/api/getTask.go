package api

import (
	"encoding/json"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/pkg/db"
	"github.com/go-chi/chi/v5"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method format, want: GET", http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")

	task, ok := db.Tasks[id]
	if !ok {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}