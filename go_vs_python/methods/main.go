package main
import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (this Point) distance(other Point) float64 {
	return math.Sqrt(math.Pow(this.x - other.x, 2) + math.Pow(this.y - other.y, 2))
}

func (this *Point) distance_better(other *Point) float64 {
	return math.Sqrt(math.Pow(this.x - other.x, 2) + math.Pow(this.y - other.y, 2))
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{2, 4}

	fmt.Println(p1.distance(p2))
	fmt.Println(p1.distance_better(&p2))
}
