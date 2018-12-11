// h 20181211
//
// Video cache & thumb

package main

import (
	"net/http"

	"../video"
)

// Health
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

// Thumb
func thumb(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	video.Thumb(w, &video.ThumbRequest{R: q.Get("r"), T: q.Get("t"), W: q.Get("w"), H: q.Get("h")})
}

// Meta
func meta(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	video.Meta(w, &video.MetaRequest{R: q.Get("r")})
}
