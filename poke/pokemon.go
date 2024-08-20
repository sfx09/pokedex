package poke

import (
	"errors"
	"fmt"
	"math/rand"
)

type pokeDex struct {
	pokemons []pokemon
}

type pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (s *State) CatchPokemon(args ...string) error {
	p := pokemon{}
	err := s.in.Query(PokemonUrl+args[0], &p)
	if err != nil {
		return errors.New("Unable to fetch results")
	}
	fmt.Printf("Throwing a ball at %v\n", p.Name)
	if catch(p) {
		fmt.Printf("%v was caught!\n", p.Name)
		s.pokedex.pokemons = append(s.pokedex.pokemons, p)
	} else {
		fmt.Printf("%v escaped.\n", p.Name)
	}
	return nil
}

func catch(p pokemon) bool {
	odds := rand.Intn(p.BaseExperience)
	return odds < 40
}
