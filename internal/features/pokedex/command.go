package pokedex

import (
	"fmt"
	"pokedex/internal/command"
	"pokedex/internal/pokedex"
)

func Register(dex *pokedex.Pokedex) map[string]command.CliCommand {
	return map[string]command.CliCommand{
		"pokedex": {
			Name:        "pokedex",
			Description: "pokedex - prints a list of all the names of the Pokemon the user has caught.",
			Callback:    func(_ *command.Config, _ string) { getPokedex(dex) },
		},
	}
}

func getPokedex(dex *pokedex.Pokedex) {
	pokemons := dex.GetAll()
	if len(pokemons) == 0 {
		fmt.Println("You have not catched any pokemon yet!!")
		return
	}

	fmt.Println("Your Pokedex:")
	for _, p := range pokemons {
		fmt.Printf(" - %s\n", p.Name)
	}
}
