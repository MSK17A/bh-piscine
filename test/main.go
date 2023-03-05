package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
	fmt.Println(piscine.Index("\\~gahTCljM]f;", "aoxsA;<1t%5f$"))
}
