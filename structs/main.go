package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func distance_better(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
    p1 := Point{x: 3, y: 4}
	p2 := Point{x: 6, y: 8}

	fmt.Println(distance(p1, p2))
	fmt.Println(distance_better(&p1, &p2))
}