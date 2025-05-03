package main

import (
	"fmt"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/pkg/api"
	
	"github.com/go-chi/chi/v5"
)


func main() {
	r := chi.NewRouter()

	r.Get("/tasks", api.GetTasks)
	r.Post("/tasks", api.PostTask)
	r.Get("/tasks/{id}", api.GetTask)
	r.Delete("/tasks/{id}", api.DeleteTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
