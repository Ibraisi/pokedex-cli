package pokedex

import (
	"sync"
)

type Pokedex struct {
	mu   sync.RWMutex
	data map[string]Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{data: make(map[string]Pokemon)}
}

func (d *Pokedex) Add(name string, p Pokemon) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[name] = p
}

func (d *Pokedex) Get(name string) (Pokemon, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	p, ok := d.data[name]
	return p, ok
}
func (d *Pokedex) GetAll() []Pokemon {
	pokemons := make([]Pokemon, 0)

	for _, v := range d.data {
		pokemons = append(pokemons, v)

	}

	return pokemons
}
