package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		// 1
		// 2
		// 3
		// 4
		// 5
	}
}
