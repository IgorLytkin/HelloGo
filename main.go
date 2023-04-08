package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	switch {
	case n > 0:
		fmt.Print("Число положительное")
	case n < 0:
		fmt.Print("Число отрицательное")
	default:
		fmt.Print("Ноль")
	}
}
