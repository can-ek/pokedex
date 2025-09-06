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
			fmt.Println("Your command was:", cleaned[0])
		}
	}
}
