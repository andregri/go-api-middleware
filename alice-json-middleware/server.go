package main

import (
	"log"
	"net/http"

	"github.com/andregri/go-api-middleware/mwr"
	"github.com/justinas/alice"
)

func main() {
	mainLogicHandler := http.HandlerFunc(mwr.MainLogic)
	chain := alice.New(mwr.FilterContentType, mwr.SetServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
