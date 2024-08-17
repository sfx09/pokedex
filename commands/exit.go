package commands

import (
	"fmt"
	"os"
)

func exitCommand(args ...string) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
