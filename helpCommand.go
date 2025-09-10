package main

import "fmt"

func commandHelp(session *sessionConfig, params ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("==========================")
	fmt.Println("Usage")
	fmt.Println("==========================")

	for _, cmd := range getCommands() {
		fmt.Println(cmd.name+":\t", cmd.description)
	}

	fmt.Println("==========================")
	return nil
}
