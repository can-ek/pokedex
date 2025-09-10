package apiclient

import (
	"time"
)

type PokeClient interface {
	GetLocationAreas(url string) (LocationAreas, error)
	GetLocationArea(name string) (LocationArea, error)
}

type pokeClientInternal struct {
	client apiClient
}

func NewPokeClient(timeout time.Duration) PokeClient {
	return &pokeClientInternal{
		client: newClient(timeout),
	}
}
