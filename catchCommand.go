package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(session *sessionConfig, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Error: Missing parameter for location area\n")
	}

	cleanParams := cleanInput(params[0])
	name := cleanParams[0]

	// get pokemon
	pokemon, err := session.pokeClient.GetPokemon(name)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// logic to catch the pokemon
	random := rand.Intn(pokemon.Experience)

	if random%2 == 0 {
		fmt.Println(pokemon.Name, "was caught!")
		session.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	// print result
	return nil
}
