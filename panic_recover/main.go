package main

import "fmt"

func main() {
    

	defer func(){
		fmt.Println("error was:",recover())
	}()

	panic("panic")

}