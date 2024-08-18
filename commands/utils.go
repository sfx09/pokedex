package commands

import (
	"encoding/json"
	"net/http"
)

func queryUrl(url string, v any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(&v)
	return err
}
