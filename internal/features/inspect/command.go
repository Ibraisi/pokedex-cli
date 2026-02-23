package inspect

import (
	"fmt"
	"pokedex/internal/command"
	"pokedex/internal/pokedex"
)

func Register(dex *pokedex.Pokedex) map[string]command.CliCommand {
	return map[string]command.CliCommand{
		"inspect": {
			Name:        "inspect",
			Description: "inspect <name> - display stats of a caught Pokemon.",
			Callback: func(cfg *command.Config, args []string) {
				if len(args) == 0 {
					fmt.Println("usage: inspect <pokename>")
					return
				}
				inspect(args[0], dex)
			},
		},
	}
}

func inspect(name string, dex *pokedex.Pokedex) {
	p, ok := dex.Get(name)
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n", p.Stats.HP)
	fmt.Printf("  -attack: %d\n", p.Stats.Attack)
	fmt.Printf("  -defense: %d\n", p.Stats.Defense)
	fmt.Printf("  -special-attack: %d\n", p.Stats.SpecialAttack)
	fmt.Printf("  -special-defense: %d\n", p.Stats.SpecialDefense)
	fmt.Printf("  -speed: %d\n", p.Stats.Speed)
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t)
	}
}
