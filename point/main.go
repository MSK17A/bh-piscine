package main

import (
	"piscine"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	piscine.PrintStr("x = ")
	piscine.PrintNbr(points.x)
	piscine.PrintStr(" ")
	piscine.PrintStr("y = ")
	piscine.PrintNbr(points.y)
	piscine.PrintStr("\n")
}
