// h 20181211
//
// Video cache & thumb

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Runner
func main() {
	// Route
	r := mux.NewRouter()
	r.HandleFunc("/health", health)
	r.HandleFunc("/thumb", thumb)
	r.HandleFunc("/meta", meta)
	http.Handle("/", r)
	// Endpoint
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	log.Println(host + ":" + port)
	// Serve
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		panic(err)
	}
}
