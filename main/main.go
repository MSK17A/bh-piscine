package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Atoi("12345"))
	fmt.Println(piscine.Atoi("0000000012345"))
	fmt.Println(piscine.Atoi("012 345"))
	fmt.Println(piscine.Atoi("Hello World!"))
	fmt.Println(piscine.Atoi("+1234"))
	fmt.Println(piscine.Atoi("-1234"))
	fmt.Println(piscine.Atoi("++1234"))
	fmt.Println(piscine.Atoi("--1234"))
	fmt.Println(piscine.Atoi("-6441424758376181640"))
	fmt.Println(piscine.Atoi("7946339792308667397"))
}
