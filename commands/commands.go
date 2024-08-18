package commands

import "fmt"

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
	mpf, mpb := newMap()
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
			execute: mpf,
		},
		"mapb": {
			Name:    "mapb",
			Desc:    "List Previous 20 Pokemon locations",
			execute: mpb,
		},
	}
}
