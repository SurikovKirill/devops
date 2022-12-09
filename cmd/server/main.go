package main

import (
	"devops/internal/handlers"
	store "devops/internal/storage"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	var m store.MemStorage
	m.Init()
	r := chi.NewRouter()
	r.Mount("/value", handlers.MetricsResource{M: &m}.Routes())
	r.Mount("/update", handlers.MetricsResourceUpdate{M: &m}.Routes())
	log.Fatal(http.ListenAndServe(":8080", r))
}
