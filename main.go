package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	switch {
	case n > 10000:
		{
			fmt.Print("Число должно быть меньше 10000 !")
			return
		}
	case n > 1000:
		{
			fmt.Print(n / 1000)
		}
	case n > 100:
		{
			fmt.Print(n / 100)
		}
	case n > 10:
		{
			fmt.Print(n / 10)
		}
	default:
		fmt.Print(n)
	}
}
