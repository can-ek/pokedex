package main

import "fmt"

func commandExplore(session *sessionConfig, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Error: Missing parameter for location area\n")
	}

	cleanParams := cleanInput(params[0])
	name := cleanParams[0]
	locationArea, err := session.pokeClient.GetLocationArea(name)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Exploring", locationArea.Name, "...")
	fmt.Println("Found Pokemon:")

	for _, encounter := range locationArea.Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
