package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := range workArray {
		_, err := fmt.Scan(&workArray[i])
		if err != nil {
			return
		}
	}
	var a, b, w1, w2 uint8
	for i := 1; i <= 3; i++ {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			return
		}
		w1 = workArray[a]
		w2 = workArray[b]
		workArray[a] = w2
		workArray[b] = w1
	}
	fmt.Println(workArray)
}
