package explore

import (
	"encoding/json"
	"fmt"
	"pokedex/internal/command"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
)

func Register(cache *pokecache.Cache) map[string]command.CliCommand {
	return map[string]command.CliCommand{
		"explore": {
			Name:        "explore",
			Description: "explore <location> - list pokemon available in a location area",
			Callback:    func(cfg *command.Config, location string) { exploreLocation(location, cache) },
		},
	}
}

func exploreLocation(location string, cache *pokecache.Cache) {
	if location == "" {
		fmt.Println("usage: explore <location>")
		return
	}

	exploreRes, err := getExploreRes(location, cache)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Exploring %s...\n", location)
	for _, pe := range exploreRes.PokemonEncounters {
		fmt.Println(pe.Pokemon.Name)
	}
}

func getExploreRes(location string, cache *pokecache.Cache) (ExploreRes, error) {
	if data, ok := cache.Get(location); ok {
		var res ExploreRes
		err := json.Unmarshal(data, &res)
		return res, err
	}

	c := pokeapi.NewPokeAPIClient()
	res, err := getPokesInLocation(c, location)
	if err != nil {
		return ExploreRes{}, err
	}

	data, err := json.Marshal(&res)
	if err != nil {
		return ExploreRes{}, err
	}
	cache.Add(location, data)
	return res, nil
}
