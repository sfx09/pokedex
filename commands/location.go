package commands

import (
	"errors"
)

type locatorAPI struct {
	prevUrl string
	nextUrl string
}

type locatorResp struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func newLocatorAPI(url string) locatorAPI {
	return locatorAPI{prevUrl: "", nextUrl: url}
}

func (l *locatorAPI) query(forwardFlag bool) (locatorResp, error) {
	resp := locatorResp{}
	var url string
	if forwardFlag {
		url = l.nextUrl
	} else {
		url = l.prevUrl
	}
	if url == "" {
		return resp, errors.New("End of results")
	}
	err := queryUrl(url, &resp)
	if err != nil {
		return resp, err
	}
	l.nextUrl = resp.Next
	l.prevUrl = resp.Previous
	return resp, err
}
