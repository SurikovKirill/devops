package handlers

import (
	"devops/internal/helpers"
	store "devops/internal/storage"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Структура хэндлера для /value
type MetricsResource struct {
	M *store.MemStorage
}

func (rs MetricsResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.list)
	r.Route("/", func(r chi.Router) {
		r.Get("/{type}/{name}", rs.get)
	})
	return r
}

const (
	html = `
	<h1>Metrics</h1>
	{{
		range .MemStorage
	}}
	{{end}}
	`
)

// Вывести все метрики
func (rs MetricsResource) list(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Parse(html))
	if err := tmpl.Execute(w, rs.M); err != nil {
		log.Fatal(err)
	}
}

// Получить метрику
func (rs MetricsResource) get(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "type")
	if t != "gauge" && t != "counter" {
		http.Error(w, "Wrong type of metric", http.StatusNotFound)
		return
	}
	v, ok := rs.M.Get(chi.URLParam(r, "name"))
	if !ok {
		http.Error(w, "Wrong name metric or doesn't exist", http.StatusNotFound)
		return
	}
	w.Header().Set("content-type", "text/plain")
	if t == "counter" {
		w.Write([]byte(fmt.Sprintf("%d", v.(helpers.Counter))))
		return
	}
	w.Write([]byte(fmt.Sprintf("%g", v.(helpers.Gauge))))
}

// Структура хэндлера для /update
type MetricsResourceUpdate struct {
	M *store.MemStorage
}

func (rs MetricsResourceUpdate) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/{type}/{name}/{val}", rs.updateMetrics)
	})
	return r
}

// Обновить значение метрики
func (rs MetricsResourceUpdate) updateMetrics(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "type")
	if t != "counter" && t != "gauge" {
		http.Error(w, "The metric doesn't exist", http.StatusNotImplemented)
		return
	}
	if t == "counter" {
		last, _ := rs.M.Get(chi.URLParam(r, "name"))
		val, err := strconv.ParseInt(chi.URLParam(r, "val"), 10, 64)
		if err != nil {
			http.Error(w, "Wrong value", http.StatusBadRequest)
			return
		}
		if last == nil {
			rs.M.Set(chi.URLParam(r, "name"), helpers.Counter(val))
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(http.StatusOK)
			return
		}
		rs.M.Set(chi.URLParam(r, "name"), helpers.Counter(val)+last.(helpers.Counter))
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		return
	}
	if t == "gauge" {
		val, err := strconv.ParseFloat(chi.URLParam(r, "val"), 64)
		if err != nil {
			http.Error(w, "Wrong value", http.StatusBadRequest)
			return
		}
		rs.M.Set(chi.URLParam(r, "name"), helpers.Gauge(val))
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		return
	}

}
