package main

import (
	"errors"
	"fmt"
)

func callbackInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided...")
	}

	pokemonName := args[0]
	pokemon, ok := c.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught the pokemon \n")
	}

	fmt.Printf("Name %s \n", pokemon.Name)
	fmt.Printf("Height : %v \n", pokemon.Height)
	fmt.Printf("Weight : %v \n", pokemon.Weight)

	fmt.Printf("Type  \n")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s \n", t.Type.Name)
	}

	fmt.Printf("Stats \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s : %v \n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Abilities  \n")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("- %s \n", ability.Ability.Name)
	}
	return nil
}
