package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("c2gmoXzd33lna nkHBUI60r68Q@ gA~-__}aGm`1b )eTS)f|{LjG->  (tP|GJ(g]5Iu NwDfgSi9K9#G, N(l(*yjNkYh/= Ux)\"/,^A[[(/P"))
}

/*
SplitWhiteSpaces("c2gmoXzd33lna nkHBUI60r68Q@ gA~-__}aGm`1b )eTS)f|{LjG->  (tP|GJ(g]5Iu NwDfgSi9K9#G, N(l(*yjNkYh/= Ux)\"/,^A[[(/P") == []string{"c2gmoXzd33lna", "nkHBUI60r68Q@", "gA~-__}aGm`1b", ")eTS)f|{LjG->", "", "(tP|GJ(g]5Iu", "NwDfgSi9K9#G,", "N(l(*yjNkYh/=", "Ux)\"/,^A[[(/P"} instead of []string{"c2gmoXzd33lna", "nkHBUI60r68Q@", "gA~-__}aGm`1b", ")eTS)f|{LjG->", "(tP|GJ(g]5Iu", "NwDfgSi9K9#G,", "N(l(*yjNkYh/=", "Ux)\"/,^A[[(/P"}

[]string{"c2gmoXzd33lna", "nkHBUI60r68Q@", "gA~-__}aGm`1b", ")eTS)f|{LjG->", "(tP|GJ(g]5Iu", "NwDfgSi9K9#G,", "N(l(*yjNkYh/=", "Ux)\"/,^A[[(/P"}
*/
