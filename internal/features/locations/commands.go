package locations

import (
	"encoding/json"
	"fmt"
	"pokedex/internal/command"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
)

func Register(cache *pokecache.Cache) map[string]command.CliCommand {
	return map[string]command.CliCommand{
		"map": {
			Name:        "map",
			Description: "displays the names of 20 location areas in the Pokemon world",
			Callback:    func(cfg *command.Config, _ []string) { getLocations(cfg, cache) },
		},
		"mapb": {
			Name:        "mapb",
			Description: "displays the previous 20 location areas in the Pokemon world",
			Callback:    func(cfg *command.Config, _ []string) { getPrevLocations(cfg, cache) },
		},
	}
}

func getLocations(config *command.Config, cache *pokecache.Cache) {
	local, err := getLocationRes(config.Next, cache)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, l := range local.Result {
		fmt.Println(l.Name)
	}
	config.Next = local.Next
	config.Prev = local.Prev
}

func getPrevLocations(config *command.Config, cache *pokecache.Cache) {
	local, err := getLocationRes(config.Prev, cache)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, l := range local.Result {
		fmt.Println(l.Name)
	}
	config.Next = local.Next
	config.Prev = local.Prev
}

func getLocationRes(url string, cache *pokecache.Cache) (LocationAreaRes, error) {
	if data, ok := cache.Get(url); ok {
		var res LocationAreaRes
		err := json.Unmarshal(data, &res)
		return res, err
	}

	c := pokeapi.NewPokeAPIClient()
	res, err := GetLocals(c, url)
	if err != nil {
		return LocationAreaRes{}, err
	}

	data, err := json.Marshal(&res)
	if err != nil {
		return LocationAreaRes{}, err
	}
	cache.Add(url, data)
	return res, nil
}
