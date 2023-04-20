package groupie

import (
	"net/http"
	"strconv"
	"text/template"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	var slideralbum [2]int
	var slidercreation [2]int
	var pays []string
	var members []int
	var err error
	data := Artists
	if err := r.ParseForm(); err != nil {
		return
	}
	for i, v := range r.PostForm {
		if i == "slidercreation" {
			slidercreation[0], _ = strconv.Atoi(v[0])
			slidercreation[1], _ = strconv.Atoi(v[1])
			data, err = DataByCreationDate(slidercreation[0], slidercreation[1], data)
			if err != nil {
				http.Redirect(w, r, "404 Error, FilterCreationDate", 404)
			}
		} else if i == "slideralbum" {
			slideralbum[0], _ = strconv.Atoi(v[0])
			slideralbum[1], _ = strconv.Atoi(v[1])
			data, err = DataByFirstAlbum(slideralbum[0], slideralbum[1], data)
			if err != nil {
				http.Redirect(w, r, "404 Error, FilterFirstalbum", 404)
			}
		} else if i == "pays" {
			pays = append(pays, v...)
			data, err = DataByLocations(pays, data)
			if err != nil {
				http.Redirect(w, r, "404 Error, FilterLocation", 404)
			}
		} else if i == "members" {
			for _, valeur := range v {
				valeurint, _ := strconv.Atoi(valeur)
				members = append(members, valeurint)
			}
			data, err = DataByNbrMembers(members, data)
			if err != nil {
				http.Redirect(w, r, "404 Error, FilterMembers", 404)
			}
		} else if i == "order" {
			code, _ := strconv.Atoi(v[0])
			data, err = DataSortedByOrder(code, data)
			if err != nil {
				http.Redirect(w, r, "404 Error, FilterMembers", 404)
			}
		}
	}
	if len(data) > 1 {
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, data)
	} else {
		tmpl := template.Must(template.ParseFiles("static/info.html"))
		tmpl.Execute(w, data[0])
	}
}
