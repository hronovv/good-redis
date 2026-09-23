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

func (s *Store) Len() int {
	return len(s.data)
}

func (s *Store) Get(key string) (string, error) {
	v, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key %s doesn't exist", key)
	}
	return v, nil
}

func (s *Store) Set(k, v string) {
	s.data[k] = v
}

func (s *Store) Delete(k string) {
	delete(s.data, k)
}

// Rename moves a value in a map from an old key to a new key
func (s *Store) Rename(old, new string) error {
	if old == new {
		return nil
	}
	v, err := s.Get(old)
	if err != nil {
		return err
	}
	s.Set(new, v)
	s.Delete(old)
	return nil
}

// Pop deletes a key from a map and returns its value.
// If the key happens to be not in the map, function returns default string value
func (s *Store) Pop(k string) (string, error) {
	v, err := s.Get(k)
	if err != nil {
		return "", err
	}
	s.Delete(k)
	return v, nil
}
