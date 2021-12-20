package main

import (
	"log"
	"net/http"

	"github.com/andregri/go-api-middleware/mwr"
)

func main() {
	mainLogicHandler := http.HandlerFunc(mwr.MainLogic)
	// The chaining may be unreadable if we add more steps!
	http.Handle("/city", mwr.FilterContentType(mwr.SetServerTimeCookie(mainLogicHandler)))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
