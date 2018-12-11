// h 20181211
//
// Video cache & thumb

package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"../video"
	"github.com/gorilla/mux"
)

// Runner
func main() {
	// Route
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	r.HandleFunc("/thumb", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		video.Thumb(w, &video.ThumbRequest{R: q.Get("r"), T: q.Get("t"), W: q.Get("w"), H: q.Get("h")})
	})
	r.HandleFunc("/meta", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		video.Meta(w, &video.MetaRequest{R: q.Get("r")})
	})
	http.Handle("/", r)
	// Serve
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(randPort())
	}
	log.Println("http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

// Endpoint
func randPort() int {
	l, _ := net.Listen("tcp", ":0")
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
