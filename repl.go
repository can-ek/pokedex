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
	callback    func(*sessionConfig) error
}

type sessionConfig struct {
	previousUrl string
	nextUrl     string
	pokeClient  apiclient.PokeClient
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
				err := cmd.callback(session)
				if err != nil {
					fmt.Printf("Error when running command %s, Error: %v", cleaned[0], err)
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
	}
}
