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
