package main

import (
	"fmt"

	pokeapi "github.com/can-ek/pokedex/pokeapi"
)

func commandMap(session *sessionConfig, params ...string) error {
	var locationAreas pokeapi.LocationAreas
	var err error

	if session.nextUrl != "" {
		locationAreas, err = session.pokeClient.GetLocationAreas(session.nextUrl)

		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		locationAreas, err = session.pokeClient.GetLocationAreas("")

		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	session.nextUrl = locationAreas.Next
	session.previousUrl = locationAreas.Previous
	return nil
}

func commandMapBack(session *sessionConfig, params ...string) error {
	if session.previousUrl != "" {
		locationAreas, err := session.pokeClient.GetLocationAreas(session.previousUrl)

		if err != nil {
			fmt.Println(err)
			return err
		}

		for _, area := range locationAreas.Results {
			fmt.Println(area.Name)
		}

		session.nextUrl = locationAreas.Next
		session.previousUrl = locationAreas.Previous
	} else {
		fmt.Println("you're on the first page")
	}
	return nil
}
