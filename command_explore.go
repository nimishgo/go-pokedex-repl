package main

import (
	"errors"
	"fmt"
)

func callbackExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location area provided...")
	}

	locationAreaName := args[0]
	locationArea, err := c.pokeApiClient.LocationAreasResp(locationAreaName)

	if err != nil {
		// log fatal exit the console with code 1
		fmt.Errorf("%s", err)
	}

	fmt.Printf("Pokemon in %s\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" ~ %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
