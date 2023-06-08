package main

import (
	"fmt"
)

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	var n = 0
	for i := 1; x < y; i++ {
		x += x * p / 100
		n++
	}
	fmt.Println(n)

}
