package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func EventLoop(eval func(...string)) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		sc.Scan()
		op := sc.Text()
		args := strings.Fields(op)
		eval(args...)
	}
}
