package main

import (
	"github.com/sfx09/pokedex/commands"
	"github.com/sfx09/pokedex/repl"
)

func main() {
	cmd := commands.NewCommandEvalutor()
	repl.EventLoop(cmd)
}
