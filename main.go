package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mbraeutig/go-plot/api"
)

func main() {
	http.HandleFunc("/", api.Index)
	http.HandleFunc("/plot", api.Plot)
	http.HandleFunc("/health", api.Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
