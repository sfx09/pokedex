package query

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

// Query and cache HTTP endpoints
type Inquisitor struct {
	interval time.Duration
	cache    map[string]Entry
	mu       *sync.Mutex
}

type Entry struct {
	createdAt time.Time
	value     []byte
}

func NewInquisitor(cacheDuration int) Inquisitor {
	i := Inquisitor{
		cache:    make(map[string]Entry),
		mu:       &sync.Mutex{},
		interval: time.Duration(cacheDuration) * time.Second,
	}
	i.reapLoop()
	return i
}

func (i *Inquisitor) Query(url string, v any) error {
	data, exists := i.Get(url)
	if exists {
		return json.Unmarshal(data, &v)
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	i.Add(url, data)
	return json.Unmarshal(data, &v)
}

func (i *Inquisitor) Get(key string) ([]byte, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	entry, exists := i.cache[key]
	if !exists {
		return nil, false
	}
	return entry.value, true
}

func (i *Inquisitor) Add(key string, val []byte) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.cache[key] = Entry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (i *Inquisitor) reap() {
	i.mu.Lock()
	defer i.mu.Unlock()
	deletionList := []string{}
	for key := range i.cache {
		expiryTime := i.cache[key].createdAt.Add(time.Duration(i.interval))
		if expiryTime.Before(time.Now()) {
			deletionList = append(deletionList, key)
		}
	}
	for _, key := range deletionList {
		delete(i.cache, key)
	}
}

func (i *Inquisitor) reapLoop() {
	t := time.NewTicker(i.interval)
	go func() {
		for range t.C {
			i.reap()
		}
	}()
}
