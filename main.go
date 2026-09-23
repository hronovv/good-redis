package main

import (
	"fmt"
)

func main() {
	store := NewStore(1)
	store.Set("lol", "32")
	v, err := store.Get("lol")
	fmt.Println(v, err)
	err = store.Set("lola", "hi!")
	if err != nil {
		fmt.Println(err)
	}
	v, err = store.Get("lola")
	fmt.Println(v, err)
}
