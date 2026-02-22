package locations

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pokedex/internal/pokeapi"
)

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaRes struct {
	Count  int            `json:"count"`
	Next   string         `json:"next"`
	Prev   string         `json:"previous"`
	Result []locationArea `json:"results"`
}

func GetLocals(client *pokeapi.PokeAPIClient, endUrl string) (LocationAreaRes, error) {
	parsedURL, err := url.Parse(endUrl)
	if err != nil {
		panic(err)
	}
	query := parsedURL.RawQuery

	res, err := client.HttpClient.Get(fmt.Sprintf("%s/location-area/?%s", client.BaseURL, query))
	if err != nil {
		return LocationAreaRes{}, err
	}
	var locationAreas LocationAreaRes
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return LocationAreaRes{}, err
	}

	return locationAreas, nil
}
