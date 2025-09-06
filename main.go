package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buffer := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if buffer.Scan() {
			input := buffer.Text()
			cleaned := cleanInput(input)

			if len(cleaned) == 0 {
				continue
			}

			if cmd, containsKey := cliCommands[cleaned[0]]; containsKey {
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
