package main

import (
	"fmt"
	g "groupie"
	"net/http"
)

var url = "localhost:8080"

func main() {
	http.HandleFunc("/", g.MainHandler)
	http.HandleFunc("/info", g.InfoHandler)
	http.HandleFunc("/filter", g.FilterHandler)
	fmt.Println("Listening at http://" + url)
	css, img := http.FileServer(http.Dir("./static")), http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", css))
	http.Handle("/assets/", http.StripPrefix("/assets/", img))
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println(err)
	}
}