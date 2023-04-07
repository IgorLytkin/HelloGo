package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}
	// здесь ваш код
	b = a * 2
	c = b + 100
	fmt.Println(c)
}
