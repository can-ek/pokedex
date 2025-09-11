package main

import (
	"fmt"
)

func commandPokedex(session *sessionConfig, params ...string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range session.pokedex {
		fmt.Println("  -", name)
	}
	return nil
}
