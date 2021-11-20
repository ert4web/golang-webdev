package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/timpum.jpeg">`)
}
