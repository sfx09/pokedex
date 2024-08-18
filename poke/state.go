package poke

import (
	"errors"
	"fmt"

	"github.com/sfx09/pokedex/query"
)

type State struct {
	in  query.Inquisitor
	loc location
}

type location struct {
	PrevUrl string
	NextUrl string
}

type locationResponse struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func NewState() State {
	return State{
		in: query.NewInquisitor(10),
		loc: location{
			PrevUrl: "",
			NextUrl: "http://pokeapi.co/api/v2/location",
		},
	}
}

func (s *State) MapForward(args ...string) error {
	r := locationResponse{}
	err := s.in.Query(s.loc.NextUrl, &r)
	if err != nil {
		return errors.New("Unable to fetch results")
	}
	s.loc.NextUrl = r.Next
	s.loc.PrevUrl = r.Previous
	for _, loc := range r.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func (s *State) MapBackward(args ...string) error {
	r := locationResponse{}
	err := s.in.Query(s.loc.PrevUrl, &r)
	if err != nil {
		return errors.New("Unable to fetch results")
	}
	s.loc.NextUrl = r.Next
	s.loc.PrevUrl = r.Previous
	for _, loc := range r.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
