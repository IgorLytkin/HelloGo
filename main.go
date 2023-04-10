package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A)
	fmt.Scan(&B)

	var s int
	s = 0
	for i := A; i <= B; i++ {
		s += i
	}
	fmt.Println(s)
}
