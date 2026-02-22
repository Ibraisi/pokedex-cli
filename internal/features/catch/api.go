package catch

import (
	"encoding/json"
	"fmt"
	"pokedex/internal/pokeapi"
)

type statInfo struct {
	Name string `json:"name"`
}

type statEntry struct {
	BaseStat int      `json:"base_stat"`
	Stat     statInfo `json:"stat"`
}

type typeInfo struct {
	Name string `json:"name"`
}

type typeEntry struct {
	Type typeInfo `json:"type"`
}

type PokemonInfoRes struct {
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	Stats          []statEntry `json:"stats"`
	Types          []typeEntry `json:"types"`
}

func getPokeBaseExperience(name string, client *pokeapi.PokeAPIClient) (PokemonInfoRes, error) {
	res, err := client.HttpClient.Get(fmt.Sprintf("%s/pokemon/%s", client.BaseURL, name))
	if err != nil {
		return PokemonInfoRes{}, err
	}
	decoder := json.NewDecoder(res.Body)
	var pokemomInfoRes PokemonInfoRes
	if err := decoder.Decode(&pokemomInfoRes); err != nil {
		return PokemonInfoRes{}, err
	}

	return pokemomInfoRes, nil
}
