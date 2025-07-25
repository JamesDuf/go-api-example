package main

import (
	"fmt"
	"net/http"

	"github.com/JamesDuf/go-api-tutorial/internal/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup logger so that when we print something out we get the line and file number
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter() //pointer to a new router (Mux type) //struct used to setup API
	handlers.Handler(r)              //call the handler function to setup the routes (add endpoints)

	fmt.Println("Starting GO API service...")

	// Start the server on port 8080
	// ListenAndServe starts an HTTP server with a given address and handler
	url := "localhost:8080"
	err := http.ListenAndServe(url, r) //start the server on port 8080 of localhost
	if err != nil {
		log.Error(err) //log error if server fails to start
	}
	log.Infof("Server started on %v", url) //log info message when server
}
