package main

import (
	"time"

	apiclient "github.com/can-ek/pokedex/pokeapi"
)

func main() {
	session := sessionConfig{
		nextUrl:     "",
		previousUrl: "",
		pokeClient:  apiclient.NewPokeClient(10 * time.Second),
		pokedex:     map[string]apiclient.Pokemon{},
	}

	startRepl(&session)
}
