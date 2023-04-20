package groupie

import (
	"errors"
	"strconv"
	"strings"
)

func DataByName(name string, data []Artist) (Artist, error) {
	for _, artist := range data {
		if artist.Name == name {
			return artist, nil
		}
	}
	return Artist{}, errors.New("not found")
}

func DataByCreationDate(start int, end int, data []Artist) ([]Artist, error) {
	var tmp []Artist
	for _, artist := range data {
		if artist.CreationDate >= start && artist.CreationDate <= end {
			tmp = append(tmp, artist)
		}
	}
	if len(tmp) != 0 {
		return tmp, nil
	} else {
		return tmp, errors.New("not found")
	}
}

func DataByFirstAlbum(start int, end int, data []Artist) ([]Artist, error) {
	var tmp []Artist
	var local int
	for _, artist := range data {
		cpt := 0
		for i, valeur := range artist.FirstAlbum {
			if valeur == 45 {
				cpt++
			}
			if cpt == 2 {
				local, _ = strconv.Atoi(artist.FirstAlbum[i+1:])
				break
			}
		}
		if local >= start && local <= end {
			tmp = append(tmp, artist)
		}
	}
	if len(tmp) != 0 {
		return tmp, nil
	} else {
		return tmp, errors.New("not found")
	}
}

func DataByNbrMembers(mem []int, data []Artist) ([]Artist, error) {
	var tmp []Artist
	for _, x := range mem {
		for _, artist := range data {
			if x == len(artist.Members) {
				tmp = append(tmp, artist)
			}
		}
	}
	if len(tmp) != 0 {
		return tmp, nil
	} else {
		return tmp, errors.New("not found")
	}
}

func DataByLocations(locator []string, data []Artist) ([]Artist, error) {
	var tmp []Artist
	for _, artist := range data {
		for _, location := range artist.Locations {
			for _, local := range locator {
				if strings.Contains(location, local) {
					tmp = append(tmp, artist)
					break
				}
			}
		}
	}
	if len(tmp) != 0 {
		return tmp, nil
	} else {
		return tmp, errors.New("not found")
	}
}

func DataById(id int) (Artist, error) {
	for _, artist := range Artists {
		if artist.Id == id {
			return artist, nil
		}
	}
	return Artist{}, errors.New("not found")
}

func DataSortedByOrder(code int, data []Artist) ([]Artist, error) {
	switch code {
	case 0: // Filter by Alphabetic order
		for index := 0; index < len(data)-1; index++ {
			minindex := index
			for nombre := index; nombre < len(data); nombre++ {
				if rune(data[minindex].Name[0]) > rune(data[nombre].Name[0]) {
					minindex = nombre
				}
			}
			data[index], data[minindex] = data[minindex], data[index]
		}
		return data, nil
	case 1: // Filter by Inversed Alphabetic order
		for index := 0; index < len(data)-1; index++ {
			minindex := index
			for nombre := index; nombre < len(data); nombre++ {
				if rune(data[minindex].Name[0]) < rune(data[nombre].Name[0]) {
					minindex = nombre
				}
			}
			data[index], data[minindex] = data[minindex], data[index]
		}
		return data, nil
	case 2: // Filter by Creation Date order
		for index := 0; index < len(data)-1; index++ {
			minindex := index
			for nombre := index; nombre < len(data); nombre++ {
				if data[minindex].CreationDate > data[nombre].CreationDate {
					minindex = nombre
				}
			}
			data[index], data[minindex] = data[minindex], data[index]
		}
		return data, nil
	case 3: // Filter by Inversed Creation Date order
		for index := 0; index < len(data)-1; index++ {
			minindex := index
			for nombre := index; nombre < len(data); nombre++ {
				if data[minindex].CreationDate < data[nombre].CreationDate {
					minindex = nombre
				}
			}
			data[index], data[minindex] = data[minindex], data[index]
		}
		return data, nil
	}
	return nil, errors.New("can't order")
}
