package main

import (
	"fmt"
)

func commandInspect(session *sessionConfig, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Error: Missing parameter for location area\n")
	}

	cleanParams := cleanInput(params[0])
	name := cleanParams[0]

	// get pokemon
	if pokemon, ok := session.pokedex[name]; ok {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")

		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Metadata.Name, stat.Value)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  -%s\n", t.Data.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
