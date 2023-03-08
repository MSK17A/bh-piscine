package main

import (
	"fmt"

	"piscine"
)

func main() {
	s := "HelloHAAhowHAAareHAAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HAA"))
}
