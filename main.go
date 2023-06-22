package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	a = array[0]
	for i := 1; i < 5; i++ {
		if a < array[i] {
			a = array[i]
		}
	}
	fmt.Println(a)
}
