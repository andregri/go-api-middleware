package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request")
		handler.ServeHTTP(rw, r)
		fmt.Println("Executing middleware after response")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic for the route "/"
	fmt.Println("Executing mainLogic()")
	w.Write([]byte("OK\n"))
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
