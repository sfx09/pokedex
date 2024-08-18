package commands

import "fmt"

func newMap() (func(...string) error, func(...string) error) {
	l := newLocatorAPI("https://pokeapi.co/api/v2/location")
	genMapIterator := func(forwardFlag bool) func(s ...string) error {
		return func(s ...string) error {
			r, err := l.query(forwardFlag)
			if err != nil {
				return err
			}
			for _, loc := range r.Results {
				fmt.Println(loc.Name)
			}
			return nil
		}
	}
	return genMapIterator(true), genMapIterator(false)
}
