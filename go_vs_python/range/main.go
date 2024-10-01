package main 
import "fmt"
func main(){
	names := []string{
		"Peter",
		"Paul",
		"Mary",
	}
	for i,name := range names {
		fmt.Println(i+1,name)
	}
}