package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		getCommand := getCommands()
		cmd, ok := getCommand[command]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := cmd.callback(c, args...)
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Print the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back..",
			callback:    callbackMapb,
		},
		"exit": {
			name:        "exit",
			description: "Turn off the pokedex",
			callback:    callBackExit,
		},
		"explore": {
			name:        "explore",
			description: "lists pokemon in location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch the pokemon with your pokeball",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "display info of caught pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "see all the caught pokemon",
			callback:    callbackPokedex,
		},
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
