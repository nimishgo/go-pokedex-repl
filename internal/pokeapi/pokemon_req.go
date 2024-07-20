package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseUrl + endpoint
	// frame our request with method , url
	// check the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		// fmt.Println("cache hits")
		pokemon := Pokemon{}
		// unmarshall the data into our go struct
		// struct are pass by value hence to modify it we need to pass it with its address
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}
	// make the http request to the client
	resp, err := c.httpclient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	// read the data from the resp body until EOF return []bytes
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	// create a go struct
	pokemon := Pokemon{}
	// unmarshall the data into our go struct
	// struct are pass by value hence to modify it we need to pass it with its address
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullUrl, data)
	return pokemon, nil
}
