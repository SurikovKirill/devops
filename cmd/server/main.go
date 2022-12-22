package main

import (
	"devops/internal/routes"
	store "devops/internal/storage"
	"log"
	"net/http"
)

func main() {
	var m store.MemStorage
	// m.Init()
	r := routes.New(&m)
	log.Fatal(http.ListenAndServe(":8080", r))
}
