package main

import "fmt"

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *Store) Set(k, v string) {
	s.data[k] = v
}

func main() {
	store := NewStore()
	store.Set("lol", "32")
	v, ok := store.Get("lol")
	fmt.Println(v, ok)
}
