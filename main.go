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
	var (
		n11, n12, n13, n21, n22, n23 int
	)
	n11 = n / 100000
	n12 = (n - n11*100000) / 10000
	n13 = (n - n11*100000 - n12*10000) / 1000

	n21 = (n - n11*100000 - n12*10000 - n13*1000) / 100
	n22 = (n - n11*100000 - n12*10000 - n13*1000 - n21*100) / 10
	n23 = n - n11*100000 - n12*10000 - n13*1000 - n21*100 - n22*10

	if n11+n12+n13 == n21+n22+n23 {
		fmt.Print("YES")

	} else {
		fmt.Print("NO")
	}
}
