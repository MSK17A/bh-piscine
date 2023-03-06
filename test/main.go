package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Capitalize("hEllo! How arE you? How+are+things+4you? h:"))
	fmt.Println(piscine.Capitalize("'/O{Q?-{R;%sCY")) // "'/O{Q-{R;%Scy"
	fmt.Println(piscine.Capitalize("=O3F\\/0-e?<7I")) // "=O3f\\/0-e?<7i"
}
