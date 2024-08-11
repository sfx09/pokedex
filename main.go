package main

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	name     string
	desc     string
	callback func() error
}

func eval(s string) {
	switch s {
	case "help":
		fmt.Println("Display help message")
	default:
		fmt.Println("Display help message")
	}
}

func eventloop(eval func(string)) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		sc.Scan()
		op := sc.Text()
		eval(op)
	}
}

func main() {
	eventloop(eval)
}
