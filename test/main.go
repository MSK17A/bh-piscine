package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("c2gmoXzd33lna nkHBUI60r68Q@ gA~-__}aGm`1b )eTS)f|{LjG->  (tP|GJ(g]5Iu NwDfgSi9K9#G, N(l(*yjNkYh/= Ux)\"/,^A[[(/P"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("j9w cOmb0[whP T0#\\`*65k:A($ \\;i$zooz%[dSP jAlc$8(k0~C3` r\"g+K}vk8Giy% fD5Uv 897FBeK B;?w/Wmk>W|f} wXIo)1%:&4)G "))
}

/*
SplitWhiteSpaces("c2gmoXzd33lna nkHBUI60r68Q@ gA~-__}aGm`1b )eTS)f|{LjG->  (tP|GJ(g]5Iu NwDfgSi9K9#G, N(l(*yjNkYh/= Ux)\"/,^A[[(/P") == []string{"c2gmoXzd33lna", "nkHBUI60r68Q@", "gA~-__}aGm`1b", ")eTS)f|{LjG->", "", "(tP|GJ(g]5Iu", "NwDfgSi9K9#G,", "N(l(*yjNkYh/=", "Ux)\"/,^A[[(/P"} instead of []string{"c2gmoXzd33lna", "nkHBUI60r68Q@", "gA~-__}aGm`1b", ")eTS)f|{LjG->", "(tP|GJ(g]5Iu", "NwDfgSi9K9#G,", "N(l(*yjNkYh/=", "Ux)\"/,^A[[(/P"}

[]string{"c2gmoXzd33lna", "nkHBUI60r68Q@", "gA~-__}aGm`1b", ")eTS)f|{LjG->", "(tP|GJ(g]5Iu", "NwDfgSi9K9#G,", "N(l(*yjNkYh/=", "Ux)\"/,^A[[(/P"}
*/

/*
SplitWhiteSpaces("j9w cOmb0[whP T0#\\`*65k:A($ \\;i$zooz%[dSP jAlc$8(k0~C3` r\"g+K}vk8Giy% fD5Uv 897FBeK B;?w/Wmk>W|f} wXIo)1%:&4)G ")

== []string{"j9w", "cOmb0[whP", "T0#\\`*65k:A($", "\\;i$zooz%[dSP", "jAlc$8(k0~C3`", "r\"g+K}vk8Giy%", "fD5Uv", "897FBeK", "B;?w/Wmk>W|f}", "wXIo)1%:&4)G", ""}

instead of []string{"j9w", "cOmb0[whP", "T0#\\`*65k:A($", "\\;i$zooz%[dSP", "jAlc$8(k0~C3`", "r\"g+K}vk8Giy%", "fD5Uv", "897FBeK", "B;?w/Wmk>W|f}", "wXIo)1%:&4)G"}
*/
