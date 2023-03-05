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
	fmt.Println(piscine.Index("\\ti^&9o<_o-_.", "&9o"))
	fmt.Println(piscine.Index("J(1cTJj[>Tdtn", "(DWK-)E4<$&2)"))
}
