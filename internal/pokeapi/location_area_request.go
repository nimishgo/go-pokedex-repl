package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}
	// frame our request with method , url
	// check the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		// fmt.Println("cache hits")
		locationAreasResp := LocationAreasResp{}
		// unmarshall the data into our go struct
		// struct are pass by value hence to modify it we need to pass it with its address
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil
	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}
	// make the http request to the client
	resp, err := c.httpclient.Do(req)

	if err != nil {
		return LocationAreasResp{}, err
	}

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	// read the data from the resp body until EOF return []bytes
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	// create a go struct
	locationAreasResp := LocationAreasResp{}
	// unmarshall the data into our go struct
	// struct are pass by value hence to modify it we need to pass it with its address
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	c.cache.Add(fullUrl, data)
	return locationAreasResp, nil
}

func (c *Client) LocationAreasResp(locationName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationName
	fullUrl := baseUrl + endpoint
	// frame our request with method , url
	// check the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		// fmt.Println("cache hits")
		locationArea := LocationArea{}
		// unmarshall the data into our go struct
		// struct are pass by value hence to modify it we need to pass it with its address
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}
	// make the http request to the client
	resp, err := c.httpclient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	// read the data from the resp body until EOF return []bytes
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	// create a go struct
	locationArea := LocationArea{}
	// unmarshall the data into our go struct
	// struct are pass by value hence to modify it we need to pass it with its address
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullUrl, data)
	return locationArea, nil
}
