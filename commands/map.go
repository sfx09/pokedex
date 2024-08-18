package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func newMap() (func(args ...string) error, func(args ...string) error) {
	it := MapIterator{nextUrl: "https://pokeapi.co/api/v2/location"}
	mapf := func(args ...string) error {
		if !it.hasNext() {
			return errors.New("Failed to fetch results")
		}
		r := it.getNext()
		for _, city := range r.Results {
			fmt.Println(city.Name)
		}
		return nil
	}
	mapb := func(args ...string) error {
		if !it.hasPrev() {
			return errors.New("Failed to fetch results")
		}
		r := it.getPrev()
		for _, city := range r.Results {
			fmt.Println(city.Name)
		}
		return nil
	}
	return mapf, mapb
}

type MapIterator struct {
	prevUrl string
	Url     string
	nextUrl string
}

type Response struct {
	Prev    string
	Next    string
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func NewMapIterator(url string) MapIterator {
	return MapIterator{
		prevUrl: "",
		Url:     "",
		nextUrl: url,
	}
}

func (m *MapIterator) hasNext() bool {
	return m.nextUrl != ""
}

func (m *MapIterator) hasPrev() bool {
	return m.prevUrl != ""
}

func (m *MapIterator) getNext() Response {
	r := queryUrl(m.nextUrl)
	m.prevUrl = m.Url
	m.Url = m.nextUrl
	m.nextUrl = r.Next
	return r
}

func (m *MapIterator) getPrev() Response {
	r := queryUrl(m.prevUrl)
	m.nextUrl = m.Url
	m.Url = m.prevUrl
	m.prevUrl = r.Prev
	return r
}

func queryUrl(url string) Response {
	r := Response{}
	resp, _ := http.Get(url)
	json.NewDecoder(resp.Body).Decode(&r)
	return r
}
