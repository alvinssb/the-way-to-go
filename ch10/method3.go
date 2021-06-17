package main

import (
	"fmt"
	"math"
)

type Point2 struct {
	x, y float64
}

func (p *Point2) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point2
	string
}

func main() {
	n := &NamedPoint{Point2{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}
