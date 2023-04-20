package groupie

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/info.html"))
	idstring, _ := strconv.Atoi(r.FormValue("overlay"))
	data, err := DataById(idstring)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "404 Error", 404)
	} else {
		tmpl.Execute(w, data)
	}
}
