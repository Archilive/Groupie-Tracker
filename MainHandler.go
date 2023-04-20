package groupie

import (
	"fmt"
	"net/http"
	"text/template"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	err := Data()
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "404 Error, MainHandler", 404)
	} else {
		tmpl.Execute(w, Artists)
	}
}
