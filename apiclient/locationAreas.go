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

func GetLocationAreas() LocationAreas {
	urlPath := "location-area"
	areas, err := get[LocationAreas](query{path: urlPath, limit: 20})

	if err != nil {
		fmt.Printf("Error getting location areas: %g\n", err)
		return LocationAreas{}
	}
	return areas
}

func GetLocationArea(url string) LocationAreas {
	areas, err := get[LocationAreas](query{url: url, limit: 20})

	if err != nil {
		fmt.Printf("Error getting location areas: %g\n", err)
		return LocationAreas{}
	}
	return areas
}
