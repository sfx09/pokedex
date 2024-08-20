package poke

import (
	"errors"
	"fmt"
)

type areaResp struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (s *State) ExploreArea(args ...string) error {
	a := areaResp{}
	err := s.in.Query(LocationUrl+args[0], &a)
	if err != nil {
		return errors.New("Unable to fetch results")
	}
	for _, encounter := range a.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
