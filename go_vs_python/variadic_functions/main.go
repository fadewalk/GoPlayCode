package main
import "fmt"

func average(numbers ...float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
    fmt.Println(average(1, 2, 3, 4, 5))
}

