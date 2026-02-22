package pokeapi

import (
	"net/http"
	"time"
)

type PokeAPIClient struct {
	BaseURL    string
	HttpClient *http.Client
}

func NewPokeAPIClient() *PokeAPIClient {
	return &PokeAPIClient{
		BaseURL: "https://pokeapi.co/api/v2/",
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
