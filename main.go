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
	var n1_1, n1_2, n1_3, n2_1, n2_2, n2_3 int
	n1_1 = n / 100000
	n1_2 = (n - n1_1*100000) / 10000
	n1_3 = (n - n1_1*100000 - n1_2*10000) / 1000

	n2_1 = (n - n1_1*100000 - n1_2*10000 - n1_3*1000) / 100
	n2_2 = (n - n1_1*100000 - n1_2*10000 - n1_3*1000 - n2_1*100) / 10
	n2_3 = n - n1_1*100000 - n1_2*10000 - n1_3*1000 - n2_1*100 - n2_2*10

	if n1_1+n1_2+n1_3 == n2_1+n2_2+n2_3 {
		fmt.Print("YES")

	} else {
		fmt.Print("NO")
	}
}
