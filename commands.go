package main

import (
	"fmt"
	"os"
)

func Eval(key string) {
	cmds := getCommands()
	cmd, exists := cmds[key]
	if !exists {
		cmds["help"].execute()
		return
	}
	cmd.execute()
}

type command struct {
	Name    string
	Desc    string
	execute func() error
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

func helpCommand() error {
	fmt.Println("Usage")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.Name, c.Desc)
	}
	return nil
}

func exitCommand() error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
