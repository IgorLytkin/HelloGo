package main

import (
	"fmt"
)

func main() {
	var N, b int
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	if N > 10000 {
		return
	}
	// здесь ваш код
	b = N % 10

	fmt.Println(b)
}
