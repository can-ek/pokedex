package apiclient

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/can-ek/pokedex/pokecache"
)

type query struct {
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
	fullUrl := q.url
	key := q.url

	if val, exists := c.cacheClient.Get(key); exists {
		return val, nil
	}

	if q.limit > 0 {
		fullUrl = fmt.Sprintf("%s?limit=%d", key, q.limit)
	}

	response, err := c.httpClient.Get(fullUrl)

	if err != nil {
		fmt.Printf("Error on GET: %s - %g\n", fullUrl, err)
		return nil, err
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("Error getting element at %s, failed with StatusCode: %d", fullUrl, response.StatusCode)
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
