package main 
import "fmt"

func main() {

	x := 5
	if x != 0 {
		fmt.Println("x is not zero")
	}

	var y []string
	// fmt.Println("y:", len(y))
	if len(y) != 0 {
		fmt.Println("y is not empty")
	} else {
		fmt.Println("y is empty")
	}

}
