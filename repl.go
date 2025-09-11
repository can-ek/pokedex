package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	apiclient "github.com/can-ek/pokedex/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*sessionConfig, ...string) error
}

type sessionConfig struct {
	previousUrl string
	nextUrl     string
	pokeClient  apiclient.PokeClient
	pokedex     map[string]apiclient.Pokemon
}

func startRepl(session *sessionConfig) {
	buffer := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if buffer.Scan() {
			input := buffer.Text()
			cleaned := cleanInput(input)

			if len(cleaned) == 0 {
				continue
			}

			if cmd, containsKey := getCommands()[cleaned[0]]; containsKey {
				err := cmd.callback(session, cleaned[1:]...)
				if err != nil {
					fmt.Printf("Error when running command %s, Error: %v\n", cleaned[0], err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	lower_case := strings.ToLower(text)
	return strings.Fields(lower_case)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Navigates to the previous page of location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon encountered in the location area specified",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon, if successful it adds the pokemon to the pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Returns information from a pokemon in the pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the pokemons that have been caught",
			callback:    commandPokedex,
		},
	}
}
