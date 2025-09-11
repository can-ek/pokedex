package apiclient

import (
	"encoding/json"
	"fmt"
)

type StatMetadata struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	Value    int          `json:"base_stat"`
	Metadata StatMetadata `json:"stat"`
}

type PokeType struct {
	Data struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	Name       string        `json:"name"`
	Experience int           `json:"base_experience"`
	Height     int           `json:"height"`
	Weight     int           `json:"weight"`
	Stats      []PokemonStat `json:"stats"`
	Types      []PokeType    `json:"types"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

func (p *pokeClientInternal) GetPokemon(name string) (Pokemon, error) {
	var result Pokemon
	url := fmt.Sprintf("%s/%s/%s", baseUrl, pokemonPath, name)

	bytes, err := p.client.get(query{url: url})
	if err != nil {
		fmt.Printf("Error getting pokemon: %g\n", err)
		return result, err
	}

	if len(bytes) == 0 {
		fmt.Println("Pokemon not found")
		return result, nil
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("Error desirializing response:", err)
		return result, err
	}

	return result, nil
}
