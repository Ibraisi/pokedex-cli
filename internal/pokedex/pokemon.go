package pokedex

type Pokemon struct {
	Name   string
	Height int
	Stats  PokemonStats
	Types  []string
	Weight int
}

type PokemonStats struct {
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}
