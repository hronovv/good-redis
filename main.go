package main

import (
	"fmt"
	"slices"
)

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k, _ := range s.data {
		keys = append(keys, k)
	}
	slices.Sort(keys) // after 1.22 sort.Strings() would call slices.Sort()
	return keys
}

func (s *Store) Get(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *Store) Set(k, v string) {
	s.data[k] = v
}

func (s *Store) Delete(k string) {
	delete(s.data, k)
}

func main() {
	store := NewStore()
	store.Set("lol", "32")
	v, ok := store.Get("lol")
	fmt.Println(v, ok)
}
