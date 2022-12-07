package main

import (
	"devops/internal/handlers"
	store "devops/internal/storage"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	m := store.New()
	m.Init()
	r := chi.NewRouter()
	// r.Get("/update", func(w http.ResponseWriter, r *http.Request) {
	// 	tmp := strings.Split(r.URL.Path, "/")
	// 	log.Println("dfsadas1111")
	// 	if tmp[len(tmp)-3] != "counter" || tmp[len(tmp)-3] != "gauge" {
	// 		http.Error(w, "The metric doesn't exist", http.StatusNotImplemented)
	// 	}
	// })
	r.Mount("/value", handlers.MetricsResource{M: m}.Routes())
	r.Mount("/update", handlers.MetricsResourceUpdate{M: m}.Routes())
	// r.Post("/update/{ip}", func(w http.ResponseWriter, r *http.Request) {
	// 	tmp := strings.Split(r.URL.Path, "/")
	// 	log.Println()
	// 	if tmp[1] != "update" {
	// 		http.Error(w, "The metric doesn't exist", http.StatusBadRequest)
	// 	}
	// 	if tmp[len(tmp)-3] != "counter" || tmp[len(tmp)-3] != "gauge" {
	// 		http.Error(w, "The metric doesn't exist", http.StatusNotImplemented)
	// 		return
	// 	}
	// })

	// r.Mount("/update/gauge/", &handlers.GaugeHandler{M: m})
	// r.Mount("/update/counter/", &handlers.GaugeHandler{M: m})
	log.Fatal(http.ListenAndServe(":8080", r))
}
