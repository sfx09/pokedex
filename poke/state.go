package poke

import (
	"github.com/sfx09/pokedex/query"
)

type State struct {
	in      query.Inquisitor
	mpState mapState
	pokedex pokeDex
}

func NewState() State {
	return State{
		in: query.NewInquisitor(1000),
		mpState: mapState{
			PrevUrl: "",
			NextUrl: LocationUrl,
		},
		pokedex: pokeDex{
			pokemons: []pokemon{},
		},
	}
}
