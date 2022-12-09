package main

import (
	"devops/internal/agntstorage"
	"log"

	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

const (
	pollInterval   = time.Duration(2 * time.Second)
	reportInterval = time.Duration(10 * time.Second)
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	var rtm runtime.MemStats
	var m agntstorage.Metrics
	m.Init()
	tickerPoll := time.NewTicker(pollInterval)
	tickerReport := time.NewTicker(reportInterval)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-tickerPoll.C:
				runtime.ReadMemStats(&rtm)
				m.Update(rtm)
				log.Println("Metrics updated")
			case <-tickerReport.C:
				m.Send()
				log.Println("Metrics has been send")
			case <-c:
				return
			}
		}
	}()
	wg.Wait()
}
