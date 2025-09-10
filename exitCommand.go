package main

import (
	"fmt"
	"os"
)

func commandExit(session *sessionConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
