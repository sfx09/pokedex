package commands

import "fmt"

func newMap() (func(...string) error, func(...string) error) {
	l := newLocatorAPI("https://pokeapi.co/api/v2/location")
	return func(s ...string) error {
			r, err := l.query(true)
			if err != nil {
				return err
			}
			for _, loc := range r.Results {
				fmt.Println(loc.Name)
			}
			return nil
		}, func(s ...string) error {
			r, err := l.query(false)
			if err != nil {
				return err
			}
			for _, loc := range r.Results {
				fmt.Println(loc.Name)
			}
			return nil
		}
}
