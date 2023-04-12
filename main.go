package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)

	var s int
	// если ввели первым не ноль
	if n != 0 {
		s = 1
		m = n
	}

	for n != 0 {
		fmt.Scan(&n)
		switch {
		case m < n: // новое число больше текущего максимума?
			{
				m = n
				s = 1
			}
		case m == n: // новое число равно текущему максимуму?
			{
				s++
			}
		}
	}
	fmt.Println(s)
}
