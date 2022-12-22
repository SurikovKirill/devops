package handlers

import (
	"devops/internal/metrics"
	store "devops/internal/storage"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type MetricsGet struct {
	M *store.MemStorage
}

func (mg *MetricsGet) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", mg.list)
	r.Route("/", func(r chi.Router) {
		r.Get("/{type}/{name}", mg.get)
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

func (mg *MetricsGet) list(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Parse(html))
	if err := tmpl.Execute(w, mg.M); err != nil {
		log.Fatal(err)
	}
}

func (mg *MetricsGet) get(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "type")
	if t != "gauge" && t != "counter" {
		http.Error(w, "Wrong type of metric", http.StatusNotFound)
		return
	}
	v, ok := mg.M.Get(chi.URLParam(r, "name"))
	if !ok {
		http.Error(w, "Wrong name metric or doesn't exist", http.StatusNotFound)
		return
	}
	w.Header().Set("content-type", "text/plain")
	if t == "counter" {
		w.Write([]byte(fmt.Sprintf("%d", v.(metrics.Counter))))
		return
	}
	w.Write([]byte(fmt.Sprintf("%g", v.(metrics.Gauge))))
}

type MetricsUpdate struct {
	M *store.MemStorage
}

func (mu *MetricsUpdate) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/{type}/{name}/{val}", mu.updateMetrics)
	})
	return r
}

func (mu *MetricsUpdate) updateMetrics(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "type")
	if t != "counter" && t != "gauge" {
		http.Error(w, "The metric doesn't exist", http.StatusNotImplemented)
		return
	}
	if t == "counter" {
		last, _ := mu.M.Get(chi.URLParam(r, "name"))
		val, err := strconv.ParseInt(chi.URLParam(r, "val"), 10, 64)
		if err != nil {
			http.Error(w, "Wrong value", http.StatusBadRequest)
			return
		}
		if last == nil {
			mu.M.Set(chi.URLParam(r, "name"), metrics.Counter(val))
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(http.StatusOK)
			return
		}
		mu.M.Set(chi.URLParam(r, "name"), metrics.Counter(val)+last.(metrics.Counter))
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
		mu.M.Set(chi.URLParam(r, "name"), metrics.Gauge(val))
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		return
	}
}
