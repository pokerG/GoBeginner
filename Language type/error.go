package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", e)
}
func Sqrt(f ErrNegativeSqrt) error {
	return f
}

func main() {
		fmt.Println(2)
		fmt.Println(-2)
}
