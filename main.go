package main

import (
	"fmt"
)

func main() {
	var n, n1, n2, n3 int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	n1 = n / 100
	n2 = (n - n1*100) / 10
	n3 = n - n1*100 - n2*10
	if n1 == n2 || n1 == n3 || n2 == n3 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
