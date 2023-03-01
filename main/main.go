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
	fmt.Println(piscine.Atoi("755+"))
	fmt.Println(piscine.Atoi("78799988999998798787987897987987797898789789078780797098098098908098098989089080987787098908989880998900"))
}
