package main

import (
	"io"
	"log"
	"net/http"
)

type rt int

func (h rt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>route1</h1>")
}

type tr int

func (h tr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>route2</h1>")
}
func main() {
	var h1 rt
	var h2 tr
	mux := http.NewServeMux()
	mux.Handle("/route1", h1)
	mux.Handle("/route2", h2)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
