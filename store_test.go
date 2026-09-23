package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnAllKeysSorted(t *testing.T) {
	store := NewStore()

	store.Set("neko", "smtfh")
	store.Set("apa", "smtah")
	store.Set("seto", "sm2th")
	store.Set("nym", "smdth")

	got := store.Keys()
	want := []string{"apa", "neko", "nym", "seto"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() -> %v, want -> %v", got, want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()

	got := store.Keys()
	if len(got) != 0 {
		t.Errorf("Keys() on empty store -> %v, want []string{}", got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore()
	store.Set("hello", "world")
	if v, ok := store.Get("hello"); !ok || v != "world" {
		t.Errorf("Get() -> (%q, %t), want -> (\"world\", true)", v, ok)
	}

	if v, ok := store.Get("missing"); ok || v != "" {
		t.Error("Get() failed")
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name       string
		data       map[string]string
		deleteKey  string
		wantLength int
	}{
		{"deletes existing key", map[string]string{"ha": "14", "b": "3"}, "ha", 1},
		{"deletes missing key", map[string]string{"aa": "44", "b": "3"}, "ha", 2},
		{"deletes key in empty map", map[string]string{}, "x", 0},
		{"deletes last key", map[string]string{"ha": "14"}, "ha", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore()
			for k, v := range tc.data {
				store.Set(k, v)
			}
			store.Delete(tc.deleteKey)
			got := store.Len()
			want := tc.wantLength
			if got != want {
				t.Errorf("after Delete(%q), len -> %d, want %d", tc.deleteKey, got, want)
			}
			if _, ok := store.Get(tc.deleteKey); ok {
				t.Errorf("after Delete(%q): key is still in the map", tc.deleteKey)
			}
		})
	}
}
