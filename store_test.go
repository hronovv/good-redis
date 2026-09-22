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
