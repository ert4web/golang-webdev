package main

import (
	"io"
	"net/http"
)

type rt string

func (h rt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>any code you want</h1>")
}
func main() {
	var h rt
	http.ListenAndServe(":8080", h)
}
