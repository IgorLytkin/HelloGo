package main

import (
	"fmt"
)

func main() {
	var d, h, m, m2 int
	_, err := fmt.Scan(&d)
	if err != nil {
		return
	}
	if d > 360 {
		return
	}
	// здесь ваш код
	m = d * 2
	h = m / 60
	m2 = m % 60
	fmt.Println("It is", h, "hours", m2, "minutes.")
}
