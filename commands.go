package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func NewCommandEvalutor() func(...string) {
	cmds := getCommands()
	return func(args ...string) {
		cmd, exists := cmds[args[0]]
		if !exists {
			cmds["help"].execute()
			return
		}
		cmd.execute(args[1:]...)
	}
}

type command struct {
	Name    string
	Desc    string
	execute func(args ...string) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			Name:    "help",
			Desc:    "Display a help message",
			execute: helpCommand,
		},
		"exit": {
			Name:    "exit",
			Desc:    "Exit the program",
			execute: exitCommand,
		},
		"map": {
			Name:    "map",
			Desc:    "List Pokemon locations",
			execute: NewMapCommand(),
		},
	}
}

func helpCommand(args ...string) error {
	fmt.Println("Usage")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.Name, c.Desc)
	}
	return nil
}

func exitCommand(args ...string) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}

func NewMapCommand() func(args ...string) error {
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
	return func(args ...string) error {
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
}
