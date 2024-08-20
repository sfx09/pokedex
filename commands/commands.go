package commands

import (
	"fmt"

	"github.com/sfx09/pokedex/poke"
)

type command struct {
	Name    string
	Desc    string
	execute func(args ...string) error
}

func NewCommandEvalutor() func(...string) {
	cmds := getCommands()
	return func(args ...string) {
		cmd, exists := cmds[args[0]]
		if !exists {
			cmds["help"].execute()
			return
		}
		err := cmd.execute(args[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]command {
	state := poke.NewState()
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
			Desc:    "List Next 20 Pokemon locations",
			execute: state.MapForward,
		},
		"mapb": {
			Name:    "mapb",
			Desc:    "List Previous 20 Pokemon locations",
			execute: state.MapBackward,
		},
		"explore": {
			Name:    "explore",
			Desc:    "List Pokemons present in a given location",
			execute: state.ExploreArea,
		},
		"catch": {
			Name:    "catch",
			Desc:    "Catch a pokemon",
			execute: state.CatchPokemon,
		},
		"inspect": {
			Name:    "inspect",
			Desc:    "Inspect a pokemon",
			execute: state.InspectPokemon,
		},
		"pokedex": {
			Name:    "pokedex",
			Desc:    "List all caught pokemons",
			execute: state.ListPokemons,
		},
	}
}
