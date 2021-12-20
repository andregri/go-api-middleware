package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic for the route "/"
	fmt.Println("Executing mainLogic()")
	w.Write([]byte("OK\n"))
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", mainLogicHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
