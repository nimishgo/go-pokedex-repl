package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(c *config) error {

	resp, err := c.pokeApiClient.ListLocationAreas(c.nextLocationAreaURL)

	if err != nil {
		// log fatal exit the console with code 1
		fmt.Errorf("%s", err)
	}
	fmt.Println("Location areas : ")
	for _, area := range resp.Results {
		fmt.Printf(" %s\n", area.Name)
	}
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous

	return nil
}

func callbackMapb(c *config) error {
	if c.prevLocationAreaURL == nil {
		return errors.New("You are on the first page \n")
	}
	resp, err := c.pokeApiClient.ListLocationAreas(c.prevLocationAreaURL)

	if err != nil {
		// print err and do exit(1) close the session
		log.Fatal(err)
	}
	fmt.Println("Location areas : ")
	for _, area := range resp.Results {
		fmt.Printf(" %s\n", area.Name)
	}
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous

	return nil
}
