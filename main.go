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
	case n%400 == 0:
		{
			fmt.Print("YES")
			return
		}
	case (n%4 == 0) && (n%100 != 0):
		{
			fmt.Print("YES")
			return
		}
	default:
		{
			fmt.Print("NO")
			return
		}
	}
}
