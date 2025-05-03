package api

import (
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/pkg/db"
	"github.com/go-chi/chi/v5"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "invalid method format, want: DELETE", http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")

	_, ok := db.Tasks[id]
	if !ok {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	delete(db.Tasks, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}