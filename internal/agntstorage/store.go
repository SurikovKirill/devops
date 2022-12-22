package agntstorage

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
)

type gauge float64

type counter int64

type Metrics struct {
	GaugeRuntimeMetrics map[string]gauge
	RandomValue         gauge
	PollCount           counter
}

func (m *Metrics) Init() {
	m.GaugeRuntimeMetrics = make(map[string]gauge)
	m.PollCount = 0
	m.GaugeRuntimeMetrics["Alloc"] = 0
	m.GaugeRuntimeMetrics["BuckHashSys"] = 0
	m.GaugeRuntimeMetrics["Frees"] = 0
	m.GaugeRuntimeMetrics["GCCPUFraction"] = 0
	m.GaugeRuntimeMetrics["GCSys"] = 0
	m.GaugeRuntimeMetrics["HeapAlloc"] = 0
	m.GaugeRuntimeMetrics["HeapIdle"] = 0
	m.GaugeRuntimeMetrics["HeapInuse"] = 0
	m.GaugeRuntimeMetrics["HeapObjects"] = 0
	m.GaugeRuntimeMetrics["HeapReleased"] = 0
	m.GaugeRuntimeMetrics["HeapSys"] = 0
	m.GaugeRuntimeMetrics["LastGC"] = 0
	m.GaugeRuntimeMetrics["Lookups"] = 0
	m.GaugeRuntimeMetrics["MCacheInuse"] = 0
	m.GaugeRuntimeMetrics["MCacheSys"] = 0
	m.GaugeRuntimeMetrics["MSpanInuse"] = 0
	m.GaugeRuntimeMetrics["MSpanSys"] = 0
	m.GaugeRuntimeMetrics["Mallocs"] = 0
	m.GaugeRuntimeMetrics["NextGC"] = 0
	m.GaugeRuntimeMetrics["NumForcedGC"] = 0
	m.GaugeRuntimeMetrics["NumGC"] = 0
	m.GaugeRuntimeMetrics["OtherSys"] = 0
	m.GaugeRuntimeMetrics["PauseTotalNs"] = 0
	m.GaugeRuntimeMetrics["StackInuse"] = 0
	m.GaugeRuntimeMetrics["StackSys"] = 0
	m.GaugeRuntimeMetrics["Sys"] = 0
	m.GaugeRuntimeMetrics["TotalAlloc"] = 0
	m.RandomValue = gauge(rand.Uint64())
}

func (m *Metrics) Update(r runtime.MemStats) {
	if m.GaugeRuntimeMetrics["Alloc"] != gauge(r.Alloc) {
		m.updateMetric("Alloc", gauge(r.Alloc))
	}
	if m.GaugeRuntimeMetrics["BuckHashSys"] != gauge(r.BuckHashSys) {
		m.updateMetric("BuckHashSys", gauge(r.BuckHashSys))
	}
	if m.GaugeRuntimeMetrics["Frees"] != gauge(r.Frees) {
		m.updateMetric("Frees", gauge(r.Frees))
	}
	if m.GaugeRuntimeMetrics["GCCPUFraction"] != gauge(r.GCCPUFraction) {
		m.updateMetric("GCCPUFraction", gauge(r.GCCPUFraction))
	}
	if m.GaugeRuntimeMetrics["GCSys"] != gauge(r.GCSys) {
		m.updateMetric("GCSys", gauge(r.GCSys))
	}
	if m.GaugeRuntimeMetrics["HeapAlloc"] != gauge(r.HeapAlloc) {
		m.updateMetric("HeapAlloc", gauge(r.HeapAlloc))
	}
	if m.GaugeRuntimeMetrics["HeapIdle"] != gauge(r.HeapIdle) {
		m.updateMetric("HeapIdle", gauge(r.HeapIdle))
	}
	if m.GaugeRuntimeMetrics["HeapInuse"] != gauge(r.HeapInuse) {
		m.updateMetric("HeapInuse", gauge(r.HeapInuse))
	}
	if m.GaugeRuntimeMetrics["HeapObjects"] != gauge(r.HeapObjects) {
		m.updateMetric("HeapObjects", gauge(r.HeapObjects))
	}
	if m.GaugeRuntimeMetrics["HeapReleased"] != gauge(r.HeapReleased) {
		m.updateMetric("HeapReleased", gauge(r.HeapReleased))
	}
	if m.GaugeRuntimeMetrics["HeapSys"] != gauge(r.HeapSys) {
		m.updateMetric("HeapSys", gauge(r.HeapSys))
	}
	if m.GaugeRuntimeMetrics["LastGC"] != gauge(r.LastGC) {
		m.updateMetric("LastGC", gauge(r.LastGC))
	}
	if m.GaugeRuntimeMetrics["Lookups"] != gauge(r.Lookups) {
		m.updateMetric("Lookups", gauge(r.Lookups))
	}
	if m.GaugeRuntimeMetrics["MCacheInuse"] != gauge(r.MCacheInuse) {
		m.updateMetric("MCacheInuse", gauge(r.MCacheInuse))
	}
	if m.GaugeRuntimeMetrics["MCacheSys"] != gauge(r.MCacheSys) {
		m.updateMetric("MCacheSys", gauge(r.MCacheSys))
	}
	if m.GaugeRuntimeMetrics["MSpanInuse"] != gauge(r.MSpanInuse) {
		m.updateMetric("MSpanInuse", gauge(r.MSpanInuse))
	}
	if m.GaugeRuntimeMetrics["MSpanSys"] != gauge(r.MSpanSys) {
		m.updateMetric("MSpanSys", gauge(r.MSpanSys))
	}
	if m.GaugeRuntimeMetrics["Mallocs"] != gauge(r.Mallocs) {
		m.updateMetric("Mallocs", gauge(r.Mallocs))
	}
	if m.GaugeRuntimeMetrics["NextGC"] != gauge(r.NextGC) {
		m.updateMetric("NextGC", gauge(r.NextGC))
	}
	if m.GaugeRuntimeMetrics["NumForcedGC"] != gauge(r.NumForcedGC) {
		m.updateMetric("NumForcedGC", gauge(r.NumForcedGC))
	}
	if m.GaugeRuntimeMetrics["NumGC"] != gauge(r.NumGC) {
		m.updateMetric("NumGC", gauge(r.NumGC))
	}
	if m.GaugeRuntimeMetrics["OtherSys"] != gauge(r.OtherSys) {
		m.updateMetric("OtherSys", gauge(r.OtherSys))
	}
	if m.GaugeRuntimeMetrics["PauseTotalNs"] != gauge(r.PauseTotalNs) {
		m.updateMetric("PauseTotalNs", gauge(r.PauseTotalNs))
	}
	if m.GaugeRuntimeMetrics["StackInuse"] != gauge(r.StackInuse) {
		m.updateMetric("StackInuse", gauge(r.StackInuse))
	}
	if m.GaugeRuntimeMetrics["StackSys"] != gauge(r.StackSys) {
		m.updateMetric("StackSys", gauge(r.StackSys))
	}
	if m.GaugeRuntimeMetrics["Sys"] != gauge(r.Sys) {
		m.updateMetric("Sys", gauge(r.Sys))
	}
	if m.GaugeRuntimeMetrics["TotalAlloc"] != gauge(r.TotalAlloc) {
		m.updateMetric("TotalAlloc", gauge(r.TotalAlloc))
	}
	m.RandomValue = gauge(rand.Uint64())
}

func (m *Metrics) updateMetric(key string, value gauge) {
	m.GaugeRuntimeMetrics[key] = value
	m.PollCount += 1
}

const url = "http://127.0.0.1:8080"

func (m *Metrics) Send() {
	for k, v := range m.GaugeRuntimeMetrics {
		if _, err := http.Post(fmt.Sprintf("%s/update/gauge/%s/%g", url, k, v), "text/plain", nil); err != nil {
			log.Println(err)
		}
	}
	if _, err := http.Post(fmt.Sprintf("%s/update/gauge/RandomValue/%g", url, m.RandomValue), "text/plain", nil); err != nil {
		log.Println(err)
	}
	if _, err := http.Post(fmt.Sprintf("%s/update/counter/PollCount/%d", url, m.PollCount), "text/plain", nil); err != nil {
		log.Println(err)
	}
}
