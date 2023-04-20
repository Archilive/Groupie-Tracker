package groupie

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations"`
	ConcertDates []string            `json:"concertDates"`
	Relations    map[string][]string `json:"relations"`
}
type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Relation struct {
	Id    int                 `json:"id"`
	Dates map[string][]string `json:"datesLocations"`
}
type Dates struct {
	Index []Date `json:"index"`
}
type Locations struct {
	Index []Location `json:"index"`
}
type Relations struct {
	Index []Relation `json:"index"`
}

var Artists []Artist

func APIRequestartists() error {
	req, erreur := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if erreur != nil {
		fmt.Println(erreur)
	}
	x, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(x, &Artists)
	return nil
}

func APIRequestlocations() (Locations, error) {
	var LocationsALL Locations
	req, erreur := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if erreur != nil {
		fmt.Println(erreur)
	}
	x, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(x, &LocationsALL)
	return LocationsALL, nil
}

func APIRequestdates() (Dates, error) {
	var DatesALL Dates
	req, erreur := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if erreur != nil {
		fmt.Println(erreur)
	}
	x, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(x, &DatesALL)
	return DatesALL, nil
}

func APIRequestrelation() (Relations, error) {
	var RelationsALL Relations
	req, erreur := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if erreur != nil {
		fmt.Println(erreur)
	}
	x, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(x, &RelationsALL)
	return RelationsALL, nil
}

func Data() error {
	err1 := APIRequestartists()
	LocationsALL, err2 := APIRequestlocations()
	DatesALL, err3 := APIRequestdates()
	RelationsALL, err4 := APIRequestrelation()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return errors.New("error in apirequest")
	}
	for i := range Artists {
		Artists[i].Locations = LocationsALL.Index[i].Locations
		Artists[i].ConcertDates = DatesALL.Index[i].Dates
		Artists[i].Relations = RelationsALL.Index[i].Dates
	}
	return nil
}
