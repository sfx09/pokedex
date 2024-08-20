package poke

import (
	"errors"
	"fmt"
	"math/rand"
)

type pokeDex []pokemon

type pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
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
		s.pokedex = append(s.pokedex, p)
	} else {
		fmt.Printf("%v escaped.\n", p.Name)
	}
	return nil
}

func catch(p pokemon) bool {
	odds := rand.Intn(p.BaseExperience)
	return odds < 200
}

func (s *State) ListPokemons(args ...string) error {
	fmt.Println("Owned Pokemons:")
	for _, p := range s.pokedex {
		fmt.Println(" - ", p.Name)
	}
	return nil
}

func (s *State) InspectPokemon(args ...string) error {
	for _, pokemon := range s.pokedex {
		if pokemon.Name == args[0] {
			fmt.Println("Name: ", pokemon.Name)
			fmt.Println("Height: ", pokemon.Height)
			fmt.Println("Weight: ", pokemon.Weight)
			fmt.Println("Stats:")
			for _, stat := range pokemon.Stats {
				fmt.Printf("  %v: %v\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types:")
			for _, ttype := range pokemon.Types {
				fmt.Printf("  - %v\n", ttype.Type.Name)
			}
		}
		return nil
	}
	fmt.Println("You have not caught this pokemon yet!")
	return nil
}
