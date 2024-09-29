package main

import "fmt"

func main() {
    number := 1

	increment := func(amount int) {
		number += amount
	}

	increment(5)
	increment(3)

	fmt.Println(number)
}

