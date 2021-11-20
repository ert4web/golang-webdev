package main

import (
	"io"
	"log"
	"net/http"
)

func h1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>route</h1>")
}
func h2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Route2</h1>")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/route1", h1)
	mux.HandleFunc("/route2", h2)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
