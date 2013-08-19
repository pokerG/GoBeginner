package main 

import(
	"os"
	"strings"
	"fmt"
	)

func main(){
	os.Setenv("FOO", "1")
	fmt.Println("FOO:",os.Getenv("FOO"))
	fmt.Println("BAR:",os.Getenv("BAR"))
	fmt.Println()
	fmt.Println(os.Environ())
	for b,e := range os.Environ(){
		fmt.Println(b)
		fmt.Println(e)
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}