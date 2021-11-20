package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type rt string

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func (h rt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
	}
	data := struct {
		Method     string
		URL        *url.URL
		Header     http.Header
		Form       url.Values
		PostForm   url.Values
		ContentLen int64
	}{
		Method:     r.Method,
		URL:        r.URL,
		Header:     r.Header,
		Form:       r.Form,
		PostForm:   r.PostForm,
		ContentLen: r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}
func main() {
	var hdl rt
	http.ListenAndServe(":8080", hdl)
}
