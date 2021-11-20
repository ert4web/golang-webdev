package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/timpum.jpeg", servePic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/timpum.jpeg">`)
}

func servePic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("timpum.jpeg")
	if err != nil {
		http.Error(w, "file not foound", http.StatusNotFound)
		return
	}
	fi, err := f.Stat()
	if err != nil {

	}
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}
