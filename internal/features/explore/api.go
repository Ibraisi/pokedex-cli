package explore

import (
	"encoding/json"
	"fmt"
	"pokedex/internal/pokeapi"
)

type pokemon struct {
	Name string `json:"name"`
}

type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type ExploreRes struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

func getPokesInLocation(client *pokeapi.PokeAPIClient, location string) (ExploreRes, error) {
	res, err := client.HttpClient.Get(fmt.Sprintf("%slocation-area/%s", client.BaseURL, location))
	if err != nil {
		return ExploreRes{}, err
	}
	var exploreRes ExploreRes
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&exploreRes); err != nil {
		return ExploreRes{}, err
	}
	return exploreRes, nil
}
