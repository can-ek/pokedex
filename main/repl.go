package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
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
				err := cmd.callback()
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
	}
}
