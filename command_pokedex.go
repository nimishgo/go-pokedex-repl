package main

import (
	"fmt"
)

func callbackPokedex(c *config, args ...string) error {
	fmt.Println("Pokemon in pokedex...")

	for _, pokemon := range c.caughtPokemon {
		fmt.Printf("- %s \n", pokemon.Name)
	}

	return nil
}
