
package main
import (
	"fmt"
	"strings"
)

func upone_list(thing []string, index int) {
	thing[index] = strings.ToUpper(thing[index])
}

func upone_map(thing map[string]string,index string){
	thing[index] = strings.ToUpper(thing[index])
}

func main() {
	list := []string{"a","b","c"}

	upone_list(list, 1)

	fmt.Println(list)


	dict := map[string]string{
		"a":"anders",
		"b":"bert",
		"c":"candy",
	}
	upone_map(dict,"b")
	fmt.Println(dict)
}
