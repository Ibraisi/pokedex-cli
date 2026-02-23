package repl

import (
	"fmt"
	"io"
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

	"github.com/chzyer/readline"
)

func StartRepl() {
	cache := pokecache.NewCache(time.Hour * 24)
	dex := dex.NewPokedex()
	config := command.NewConfig()
	commands := buildCommands(cache, dex)

	rl, err := readline.New("Pokedex> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			break
		}
		words := cleanUpInput(line)
		if len(words) == 0 {
			continue
		}

		v, ok := commands[words[0]]
		if ok {
			v.Callback(config, words[1:])
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
		Callback:    func(cfg *command.Config, _ []string) { commandHelp(commands) },
	}
	return commands
}

func commandExit(config *command.Config, _ []string) {
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
