package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided...")
	}

	pokemonName := args[0]
	pokemon, err := c.pokeApiClient.GetPokemon(pokemonName)

	if err != nil {
		// log fatal exit the console with code 1
		fmt.Errorf("%s", err)
	}
	threshold := 70
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing pokeball at %s ... \n", pokemonName)
	if randNum > threshold {
		return fmt.Errorf("%s escaped \n", pokemonName)
	}
	c.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught\n", pokemonName)

	return nil
}
