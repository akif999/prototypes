package main

import (
	"fmt"

	"github.com/stretchr/stew/slice"
)

func main() {
	sl := []string{"foo", "bar", "baz"}
	fmt.Println(slice.Contains(sl, "foo"))
	fmt.Println(slice.Contains(sl, "poo"))
}
