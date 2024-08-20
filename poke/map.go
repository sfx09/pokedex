package poke

import (
	"errors"
	"fmt"
)

type mapState struct {
	PrevUrl string
	NextUrl string
}

type mapResp struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func (s *State) MapForward(args ...string) error {
	r := mapResp{}
	err := s.in.Query(s.mpState.NextUrl, &r)
	if err != nil {
		return errors.New("Unable to fetch results")
	}
	s.mpState.NextUrl = r.Next
	s.mpState.PrevUrl = r.Previous
	for _, loc := range r.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func (s *State) MapBackward(args ...string) error {
	r := mapResp{}
	err := s.in.Query(s.mpState.PrevUrl, &r)
	if err != nil {
		return errors.New("Unable to fetch results")
	}
	s.mpState.NextUrl = r.Next
	s.mpState.PrevUrl = r.Previous
	for _, loc := range r.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
