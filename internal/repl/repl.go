package repl

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"pokedex/internal/command"
	"pokedex/internal/features/catch"
	"pokedex/internal/features/explore"
	"pokedex/internal/features/inspect"
	"pokedex/internal/features/locations"
	"pokedex/internal/features/pokedex"
	"pokedex/internal/pokecache"
	dex "pokedex/internal/pokedex"
	"strings"
	"time"
)

func StartRepl() {
	cache := pokecache.NewCache(time.Hour * 24)
	dex := dex.NewPokedex()
	config := command.NewConfig()
	reader := bufio.NewScanner(os.Stdin)
	commands := buildCommands(cache, dex)

	for {
		fmt.Print("Pokedex> ")
		if !reader.Scan() {
			break
		}

		words := cleanUpInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		v, ok := commands[words[0]]
		if ok {
			location := ""
			if len(words) > 1 {
				location = words[1]
			}
			v.Callback(config, location)
		} else {
			fmt.Println("unknown command")
		}
	}
}

func buildCommands(cache *pokecache.Cache, dex *dex.Pokedex) map[string]command.CliCommand {
	commands := map[string]command.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}

	maps.Copy(commands, locations.Register(cache))
	maps.Copy(commands, explore.Register(cache))
	maps.Copy(commands, catch.Register(cache, dex))
	maps.Copy(commands, inspect.Register(dex))
	maps.Copy(commands, pokedex.Register(dex))

	commands["help"] = command.CliCommand{
		Name:        "help",
		Description: "Print help message",
		Callback:    func(cfg *command.Config, _ string) { commandHelp(commands) },
	}
	return commands
}

func commandExit(config *command.Config, _ string) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp(commands map[string]command.CliCommand) {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	fmt.Println()
}

func cleanUpInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
