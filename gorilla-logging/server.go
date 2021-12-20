package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Start processing request")
	w.Write([]byte("OK"))
	log.Println("Finished processing request")
}

func main() {
	r := mux.NewRouter()

	// Attach mainLogic handler to the / route
	r.HandleFunc("/", mainLogic)

	// Wrap the router with the logging middleware
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
