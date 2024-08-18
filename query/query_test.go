package query

import (
	"testing"
	"time"
)

func TestInquisitor(t *testing.T) {
	in := NewInquisitor(10)
	err := in.Query("http://pokeapi.co/api/v2/location?offset=20", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		start := time.Now()
		err := in.Query("http://pokeapi.co/api/v2/location?offset=20", nil)
		if err != nil {
			t.Fatal(err)
		}
		end := time.Now()
		if end.Sub(start) > time.Millisecond {
			t.Fatal("Expected query to be cached")
		}
	}
}

func TestInquisitorCacheFail(t *testing.T) {
	in := NewInquisitor(1)
	err := in.Query("http://pokeapi.co/api/v2/location?offset=20", nil)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Duration(3) * time.Second)
	start := time.Now()
	err = in.Query("http://pokeapi.co/api/v2/location?offset=20", nil)
	if err != nil {
		t.Fatal(err)
	}
	end := time.Now()

	if end.Sub(start) < time.Millisecond {
		t.Fatal("Expected query to not be cached", end.Sub(start))
	}
}
