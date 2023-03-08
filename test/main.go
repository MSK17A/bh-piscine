package main

import (
	"fmt"

	"piscine"
)

func main() {
	s := "Baasaa,,Taasa,,Jaaakaaa?"
	fmt.Printf("%#v\n", piscine.Split(s, ",,"))
}
