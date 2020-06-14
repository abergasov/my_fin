package main

import (
	"html/template"
	"net/http"
)

type myFinPage struct {
	Domain string
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, myFinPage{
			Domain: "",
		})
	})
	http.ListenAndServe(":8080", nil)
}
