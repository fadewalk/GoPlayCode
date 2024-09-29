package main
import (
	"fmt"
	"time"
)

func main(){
	t0 := time.Now()
	time.Sleep(time.Second*2)
	elapsed := time.Since(t0)
	fmt.Printf("Took %s\n", elapsed)
}