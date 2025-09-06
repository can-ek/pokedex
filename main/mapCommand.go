package main

import (
	"fmt"

	apiclient "github.com/can-ek/pokedex/apiclient"
)

// Track the current state of the map navigation
var currentMapResults apiclient.LocationAreas

func commandMap() error {
	locationAreas := apiclient.GetLocationAreas()

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	currentMapResults = locationAreas
	return nil
}
