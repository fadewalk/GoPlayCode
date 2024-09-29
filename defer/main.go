package main
import "os"

func main() {
	f,_ := os.Open("./defer.py")
	defer f.Close() // defer 关键字会将 f.Close() 放入栈中，在函数结束时执行
}

