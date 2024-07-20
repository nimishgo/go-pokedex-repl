package main

import (
	"time"

	"github.com/nimishgo/go-pokedex-repl/internal/pokeapi"
)

type config struct {
	pokeApiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(time.Hour),
	}
	for {
		startRepl(&cfg)
	}
}
