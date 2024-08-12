package main

import (
	"fmt"
	"os"
)

func Eval(args ...string) {
	cmds := getCommands()
	cmd, exists := cmds[args[0]]
	if !exists {
		cmds["help"].execute()
		return
	}
	cmd.execute(args[1:]...)
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
