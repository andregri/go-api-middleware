package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Check if it is a method post
	if r.Method == "POST" {
		// Decode JSON
		var cityPayload city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&cityPayload)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		// Instead of creating a resource using received data,
		// we print data to stdout
		fmt.Printf("Got %s city with area of %d sq miles!\n",
			cityPayload.Name, cityPayload.Area)

		// Set the response status to OK
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		// Say that the method is not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
	}
}

func main() {
	http.HandleFunc("/city", mainLogic)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
