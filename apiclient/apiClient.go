package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl string = "https://pokeapi.co/api/v2/"

type query struct {
	path  string
	limit int
}

func get[T any](q query) (T, error) {
	var result T

	fullUrl := baseUrl + q.path
	if q.limit > 0 {
		fullUrl = fmt.Sprintf("%s?limit=%d", fullUrl, q.limit)
	}

	response, err := http.DefaultClient.Get(fullUrl)

	if err != nil {
		fmt.Printf("Error on GET: %s - %g\n", fullUrl, err)
		return result, err
	}

	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return result, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("Error desirializing response:", err)
		return result, err
	}

	return result, nil
}
