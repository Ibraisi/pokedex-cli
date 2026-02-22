# Pokedex

A Pokedex in a command-line REPL. This project is part of the backend track from [boot.dev](https://www.boot.dev/courses/build-pokedex-cli-golang).

Data is fetched from [PokéAPI](https://pokeapi.co/) using GET requests, with local caching to avoid redundant calls.

---

## What I learned from this project

- How to parse JSON in Go
- Making HTTP requests in Go
- Building a CLI tool that interacts with a backend server
- Local Go development and tooling
- Caching and how it improves performance
- Slice-based package architecture in Go

---

## REPL commands

| Command | Usage | Description |
|---|---|---|
| `map` | `map` | Show the next 20 location areas |
| `mapb` | `mapb` | Show the previous 20 location areas |
| `explore` | `explore <area>` | List Pokemon found in a location area |
| `catch` | `catch <name>` | Try to catch a Pokemon and add it to your Pokedex |
| `inspect` | `inspect <name>` | Show the stats of a caught Pokemon |
| `pokedex` | `pokedex` | List all Pokemon you have caught |
| `help` | `help` | Print available commands |
| `exit` | `exit` | Exit the Pokedex |

---

## Running the project

Requires [Go](https://go.dev/) and [just](https://github.com/casey/just).

```sh
just run
```
You can also use `just --choose` to interactively cycle through and pick a recipe to run.

### REPL crashes with `signal: killed`

If the REPL crashes on startup with `signal: killed`, clear the Go build cache and try again:

```sh
just clean-cache
just run
```

---

## In progress features

**Suggested by the project author**

- [ ] Support the up arrow key to cycle through previous commands
- [ ] Simulate battles between Pokemon
- [ ] Keep Pokemon in a party and allow them to level up
- [ ] Persist the Pokedex to disk so progress is saved between sessions
- [ ] Use the PokeAPI to make exploration more interesting — for example, get choices of areas instead of typing names manually
- [ ] Support different Pokeball types (Pokeball, Great Ball, Ultra Ball) with different catch rates
- [ ] Add more unit tests

**My own ideas**

- [ ] Turn this into a full terminal UI application
- [ ] Explore different architectures for the project — currently using a slice-based package layout
- [ ] Add a Dockerfile so it can run without needing Go installed locally
