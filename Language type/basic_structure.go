package main

import "fmt"

const (
	Byte = 1 << (10 * iota)
	KB
	MB
	GB
)

func main() {

	fmt.Println(Byte)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

}
