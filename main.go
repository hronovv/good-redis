package main

import (
	"fmt"
)

func main() {
	store := NewStore()
	store.Set("lol", "32")
	v, ok := store.Get("lol")
	fmt.Println(v, ok)
}
