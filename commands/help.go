package commands

import "fmt"

func helpCommand(args ...string) error {
	fmt.Println("Usage")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.Name, c.Desc)
	}
	return nil
}
