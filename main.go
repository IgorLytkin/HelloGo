package main

import (
	"fmt"
)

func main() {
	var n int

	for 1 == 1 {
		fmt.Scan(&n)
		if n < 10 {
			continue
		}
		if n > 100 {
			break
		}
		fmt.Println(n)
	}
}
