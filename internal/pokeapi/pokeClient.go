package apiclient

import (
	"time"
)

type PokeClient interface {
	GetLocationAreas() (LocationAreas, error)
	GetLocationArea(url string) (LocationAreas, error)
}

type pokeClientInternal struct {
	client apiClient
}

func NewPokeClient(timeout time.Duration) PokeClient {
	return &pokeClientInternal{
		client: newClient(timeout),
	}
}
