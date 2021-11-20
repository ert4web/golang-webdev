package main

import (
	"log"
	"net/http"
	"text/template"
)

type rt string

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var h rt
	http.ListenAndServe(":8080", h)
}

// To use Form and PostForm you must use ParseForm()
// Form: contains values from both url values(query params) and body frot POST/PUT/PATCH
// PostForm: contains only body params from POST/PUT/PATCH
func (h rt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}
