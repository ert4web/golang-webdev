package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")

	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<form method="POST">
	<input type="text" name="q"><br/>
	<input type="submit">
	</form><br>`+v)
}
