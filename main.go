package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)

	var s int
	s = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		if (m%8 == 0) && (m >= 10 && m < 100) {
			s += m
		}
	}
	fmt.Println(s)
}
