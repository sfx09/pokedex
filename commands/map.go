package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func newMap() (func(args ...string) error, func(args ...string) error) {
	type Response struct {
		Previous string
		Next     string
		Results  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"results"`
	}
	resp := Response{}
	url := "https://pokeapi.co/api/v2/location/"
	mapf := func(args ...string) error {
		r, err := http.Get(url)
		if err != nil {
			return err
		}
		err = json.NewDecoder(r.Body).Decode(&resp)
		if err != nil {
			return err
		}
		for _, city := range resp.Results {
			fmt.Println(city.Name)
		}
		url = resp.Next
		return nil
	}
	return mapf, mapf
}
