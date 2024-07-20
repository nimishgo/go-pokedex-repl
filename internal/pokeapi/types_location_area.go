package pokeapi

import (
	"net/http"
	"time"

	"github.com/nimishgo/go-pokedex-repl/internal/pokecache"
)

// next and prev are sometimes null
// we can represent it with *string which can be null
const baseUrl = "https://pokeapi.co/api/v2"

// struct to define a client
type Client struct {
	cache      pokecache.Cache
	httpclient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpclient: http.Client{
			// to prevent hanging so we don't get stuck in a network loop
			Timeout: time.Minute,
		},
	}
}
