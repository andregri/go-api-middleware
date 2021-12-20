package mwr

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Struct to store JSON data
type city struct {
	Name string
	Area uint64
}

func MainLogic(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Got %s city with area of %d sq miles!\n",
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

// Middleware to check if the content type is JSON
func FilterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Check the content, then handle the request.
		// This middleware acts before the mainLogic
		log.Println("Currently in the check content type middleware")

		// Filter requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			rw.WriteHeader(http.StatusUnsupportedMediaType)
			rw.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
		} else {
			handler.ServeHTTP(rw, r)
		}
	})
}

// Middleware to add timestamp to response cookie
func SetServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Do business logic, then apply the server time to cookie.
		// This middleware acts after mainLogic
		handler.ServeHTTP(rw, r)

		// Create cookie
		cookie := http.Cookie{
			Name:  "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}

		// Set the cookie
		http.SetCookie(rw, &cookie)

		// Print where we are
		log.Println("Currently in the set server time middleware")
	})
}
