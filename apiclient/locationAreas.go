package apiclient

import (
	"fmt"
)

type LocationArea struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type LocationAreas struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type RequestParameters struct {
	Url string
	Id  int
}

func GetLocationAreas() LocationAreas {
	urlPath := "location-area"
	areas, err := get[LocationAreas](query{path: urlPath, limit: 20})

	if err != nil {
		fmt.Printf("Error getting location areas: %g\n", err)
		return LocationAreas{}
	}
	return areas
}

func GetLocationArea(path string) LocationAreas {
	areas, err := get[LocationAreas](query{path: path, limit: 20})

	if err != nil {
		fmt.Printf("Error getting location areas: %g\n", err)
		return LocationAreas{}
	}
	return areas
}
