package main

import (
	"fmt"

	apiclient "github.com/can-ek/pokedex/apiclient"
)

func commandMap(navigation *navigationProps) error {
	var locationAreas apiclient.LocationAreas
	var navProps navigationProps

	if navigation.nextUrl != "" {
		locationAreas = apiclient.GetLocationArea(navigation.nextUrl)
	} else {
		locationAreas = apiclient.GetLocationAreas()
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	navProps.nextUrl = locationAreas.Next
	navProps.previousUrl = locationAreas.Previous
	*navigation = navProps
	return nil
}

func commandMapBack(navigation *navigationProps) error {
	if navigation.previousUrl != "" {
		locationAreas := apiclient.GetLocationArea(navigation.previousUrl)

		for _, area := range locationAreas.Results {
			fmt.Println(area.Name)
		}

		navProps := navigationProps{nextUrl: locationAreas.Next, previousUrl: locationAreas.Previous}
		*navigation = navProps
	} else {
		fmt.Println("you're on the first page")
	}
	return nil
}
