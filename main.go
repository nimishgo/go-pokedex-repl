package main

import (
	"time"

	"github.com/nimishgo/go-pokedex-repl/internal/pokeapi"
)

type config struct {
	pokeApiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	for {
		startRepl(&cfg)
	}
}
