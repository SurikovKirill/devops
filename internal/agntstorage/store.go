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
		m.GaugeRuntimeMetrics["Alloc"] = gauge(r.Alloc)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["BuckHashSys"] != gauge(r.BuckHashSys) {
		m.GaugeRuntimeMetrics["BuckHashSys"] = gauge(r.BuckHashSys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["Frees"] != gauge(r.Frees) {
		m.GaugeRuntimeMetrics["Frees"] = gauge(r.Frees)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["GCCPUFraction"] != gauge(r.GCCPUFraction) {
		m.GaugeRuntimeMetrics["GCCPUFraction"] = gauge(r.GCCPUFraction)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["GCSys"] != gauge(r.GCSys) {
		m.GaugeRuntimeMetrics["GCSys"] = gauge(r.GCSys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["HeapAlloc"] != gauge(r.HeapAlloc) {
		m.GaugeRuntimeMetrics["HeapAlloc"] = gauge(r.HeapAlloc)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["HeapIdle"] != gauge(r.HeapIdle) {
		m.GaugeRuntimeMetrics["HeapIdle"] = gauge(r.HeapIdle)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["HeapInuse"] != gauge(r.HeapInuse) {
		m.GaugeRuntimeMetrics["HeapInuse"] = gauge(r.HeapInuse)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["HeapObjects"] != gauge(r.HeapObjects) {
		m.GaugeRuntimeMetrics["HeapObjects"] = gauge(r.HeapObjects)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["HeapReleased"] != gauge(r.HeapReleased) {
		m.GaugeRuntimeMetrics["HeapReleased"] = gauge(r.HeapReleased)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["HeapSys"] != gauge(r.HeapSys) {
		m.GaugeRuntimeMetrics["HeapSys"] = gauge(r.HeapSys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["LastGC"] != gauge(r.LastGC) {
		m.GaugeRuntimeMetrics["LastGC"] = gauge(r.LastGC)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["Lookups"] != gauge(r.Lookups) {
		m.GaugeRuntimeMetrics["Lookups"] = gauge(r.Lookups)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["MCacheInuse"] != gauge(r.MCacheInuse) {
		m.GaugeRuntimeMetrics["MCacheInuse"] = gauge(r.MCacheInuse)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["MCacheSys"] != gauge(r.MCacheSys) {
		m.GaugeRuntimeMetrics["MCacheSys"] = gauge(r.MCacheSys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["MSpanInuse"] != gauge(r.MSpanInuse) {
		m.GaugeRuntimeMetrics["MSpanInuse"] = gauge(r.MSpanInuse)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["MSpanSys"] != gauge(r.MSpanSys) {
		m.GaugeRuntimeMetrics["MSpanSys"] = gauge(r.MSpanSys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["Mallocs"] != gauge(r.Mallocs) {
		m.GaugeRuntimeMetrics["Mallocs"] = gauge(r.Mallocs)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["NextGC"] != gauge(r.NextGC) {
		m.GaugeRuntimeMetrics["NextGC"] = gauge(r.NextGC)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["NumForcedGC"] != gauge(r.NumForcedGC) {
		m.GaugeRuntimeMetrics["NumForcedGC"] = gauge(r.NumForcedGC)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["NumGC"] != gauge(r.NumGC) {
		m.GaugeRuntimeMetrics["NumGC"] = gauge(r.NumGC)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["OtherSys"] != gauge(r.OtherSys) {
		m.GaugeRuntimeMetrics["OtherSys"] = gauge(r.OtherSys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["PauseTotalNs"] != gauge(r.PauseTotalNs) {
		m.GaugeRuntimeMetrics["PauseTotalNs"] = gauge(r.PauseTotalNs)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["StackInuse"] != gauge(r.StackInuse) {
		m.GaugeRuntimeMetrics["StackInuse"] = gauge(r.StackInuse)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["StackSys"] != gauge(r.StackSys) {
		m.GaugeRuntimeMetrics["StackSys"] = gauge(r.StackSys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["Sys"] != gauge(r.Sys) {
		m.GaugeRuntimeMetrics["Sys"] = gauge(r.Sys)
		m.PollCount += 1
	}
	if m.GaugeRuntimeMetrics["TotalAlloc"] != gauge(r.TotalAlloc) {
		m.GaugeRuntimeMetrics["TotalAlloc"] = gauge(r.TotalAlloc)
		m.PollCount += 1
	}
	m.RandomValue = gauge(rand.Uint64())
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
