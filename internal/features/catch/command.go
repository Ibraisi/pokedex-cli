package catch

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"pokedex/internal/command"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
	"pokedex/internal/pokedex"
)

func Register(cache *pokecache.Cache, dex *pokedex.Pokedex) map[string]command.CliCommand {
	return map[string]command.CliCommand{
		"catch": {
			Name:        "catch",
			Description: "catch <name> - catching Pokemon adds them to the user's Pokedex.",
			Callback:    func(cfg *command.Config, name string) { catch(name, cache, dex) },
		},
	}
}

func catch(name string, cache *pokecache.Cache, dex *pokedex.Pokedex) {
	if name == "" {
		fmt.Println("usage: catch <pokename>")
		return
	}
	pokeInfoRes, err := getPokeInfoRes(name, cache)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	hitDamge := rand.IntN(pokeInfoRes.BaseExperience + 100)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if hitDamge > pokeInfoRes.BaseExperience {
		fmt.Printf("%s was caught!\n", name)
		dex.Add(name, toPokemon(pokeInfoRes))
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
}

func toPokemon(res PokemonInfoRes) pokedex.Pokemon {
	var stats pokedex.PokemonStats
	for _, s := range res.Stats {
		switch s.Stat.Name {
		case "hp":
			stats.HP = s.BaseStat
		case "attack":
			stats.Attack = s.BaseStat
		case "defense":
			stats.Defense = s.BaseStat
		case "special-attack":
			stats.SpecialAttack = s.BaseStat
		case "special-defense":
			stats.SpecialDefense = s.BaseStat
		case "speed":
			stats.Speed = s.BaseStat
		}
	}

	types := make([]string, 0, len(res.Types))
	for _, t := range res.Types {
		types = append(types, t.Type.Name)
	}

	return pokedex.Pokemon{
		Name:   res.Name,
		Height: res.Height,
		Weight: res.Weight,
		Stats:  stats,
		Types:  types,
	}
}

func getPokeInfoRes(name string, cache *pokecache.Cache) (PokemonInfoRes, error) {
	if data, ok := cache.Get(name); ok {
		var res PokemonInfoRes
		err := json.Unmarshal(data, &res)
		return res, err
	}

	c := pokeapi.NewPokeAPIClient()
	pokemonInfoRes, err := getPokeBaseExperience(name, c)
	if err != nil {
		return PokemonInfoRes{}, err
	}
	data, err := json.Marshal(pokemonInfoRes)
	if err != nil {
		return PokemonInfoRes{}, err
	}
	cache.Add(name, data)

	return pokemonInfoRes, nil
}
