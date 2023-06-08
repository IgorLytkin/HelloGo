package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	var n, m int
	n = x
	m = y
	for i := 10000; i >= 1; i /= 10 {
		if n/i == 0 { // если текущая цифра ноль
			continue
		}
		m = y // возвращаем значение y
		for j := 10000; j >= 1; j /= 10 {
			if m/j == 0 {
				continue
			}
			if n/i == m/j {
				fmt.Print(n/i, " ")
			}
			m -= m / j * j
		}
		n -= n / i * i
	}
}
