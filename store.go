package main

import (
	"errors"
	"fmt"
	"slices"
)

var ErrEmptyKey error = errors.New("Empty key is not allowed")
var ErrSizeLimit error = errors.New("You have reached the limit of the keys")

type Store struct {
	maxSize int
	data    map[string]string
}

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
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
	if key == "" {
		return "", ErrEmptyKey
	}
	v, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("Key %s doesn't exist", key)
	}
	return v, nil
}

func (s *Store) Set(k, v string) error {
	if k == "" {
		return ErrEmptyKey
	}

	if _, exists := s.data[k]; s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q) -> %w", k, ErrSizeLimit)
	}

	s.data[k] = v
	return nil
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
