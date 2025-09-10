package apiclient

import (
	"encoding/json"
	"fmt"
)

type LocationAreaRef struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type LocationAreas struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []LocationAreaRef `json:"results"`
}

type LocationArea struct {
	Encounters []PokemonEncounter `json:"pokemon_encounters"`
	Name       string             `json:"name"`
}

func (p *pokeClientInternal) GetLocationAreas(url string) (LocationAreas, error) {
	var result LocationAreas

	if url == "" {
		url = fmt.Sprintf("%s/%s", baseUrl, locationAreaPath)
	}

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

func (p *pokeClientInternal) GetLocationArea(name string) (LocationArea, error) {
	var result LocationArea
	url := fmt.Sprintf("%s/%s/%s", baseUrl, locationAreaPath, name)

	bytes, err := p.client.get(query{url: url})
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
