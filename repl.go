package main

import (
	"bufio"
	"fmt"
	"os"
)

func eventloop(eval func(string)) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		sc.Scan()
		op := sc.Text()
		eval(op)
	}
}
