package main

import "fmt"

func callbackHelp(c *config) error {
	fmt.Println(" Welcome to pokedex help menu")
	fmt.Println(" Here are your available command")
	allCommands := getCommands()
	for _, v := range allCommands {
		fmt.Printf(" %s : %s \n", v.name, v.description)
	}

	return nil
}
