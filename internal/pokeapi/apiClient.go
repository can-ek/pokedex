package apiclient

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/can-ek/pokedex/pokecache"
)

const baseUrl string = "https://pokeapi.co/api/v2/"

type query struct {
	path  string
	url   string
	limit int
}

type apiClient interface {
	get(q query) ([]byte, error)
}

type client struct {
	cacheClient pokecache.CacheClient
	httpClient  *http.Client
}

func newClient(timeout time.Duration) apiClient {
	return &client{
		cacheClient: pokecache.NewCacheClient(timeout),
		httpClient:  &http.Client{Timeout: timeout},
	}
}

func (c *client) get(q query) ([]byte, error) {
	var fullUrl string
	var key string

	if q.url != "" {
		key = q.url
	} else {
		key = baseUrl + q.path
	}

	if val, exists := c.cacheClient.Get(key); exists {
		return val, nil
	}

	if q.limit > 0 {
		fullUrl = fmt.Sprintf("%s?limit=%d", key, q.limit)
	}

	response, err := http.DefaultClient.Get(fullUrl)

	if err != nil {
		fmt.Printf("Error on GET: %s - %g\n", fullUrl, err)
		return nil, err
	}

	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	err = c.cacheClient.Add(key, bytes)
	if err != nil {
		fmt.Println("Error adding", key, "to cache, continue...")
	}

	return bytes, nil
}
