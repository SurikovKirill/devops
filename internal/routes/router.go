package routes

import (
	"devops/internal/handlers"
	store "devops/internal/storage"

	"github.com/go-chi/chi"
)

func New(st *store.MemStorage) *chi.Mux {
	r := chi.NewRouter()
	mg := handlers.MetricsGet{M: st}
	mu := handlers.MetricsUpdate{M: st}
	r.Mount("/value", mg.Routes())
	r.Mount("/update", mu.Routes())
	return r
}
