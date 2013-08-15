package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	n := make(map[string]int)
	for k, v := range m {
		n[v] = k
	}
	fmt.Println(m)
	fmt.Println(n)
}
