package apiclient

import (
	"encoding/json"
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

func (p *pokeClientInternal) GetLocationAreas() (LocationAreas, error) {
	urlPath := "location-area"
	var result LocationAreas
	bytes, err := p.client.get(query{path: urlPath, limit: 20})

	if err != nil {
		fmt.Printf("Error getting location areas: %g\n", err)
		return result, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("Error desirializing response:", err)
		return result, err
	}

	return result, nil
}

func (p *pokeClientInternal) GetLocationArea(url string) (LocationAreas, error) {
	var result LocationAreas

	bytes, err := p.client.get(query{url: url, limit: 20})
	if err != nil {
		fmt.Printf("Error getting location areas: %g\n", err)
		return result, nil
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("Error desirializing response:", err)
		return result, err
	}

	return result, nil
}
