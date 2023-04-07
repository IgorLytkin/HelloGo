package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	var y float64 = math.Pi
	var max, min int
	var x2 int = 10
	var c string = "Hello World!"
	var z float64 = 1.045
	var a = 12
	var hello = "Hello"

	fmt.Println("Hello, Go!", x, y, max, min, x2, z, c, a, hello)

	var symbol int32 = 'c'
	fmt.Println(string(symbol))

	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)

	fmt.Print("hello, world")
	fmt.Print("hello, world")
	// вывод будет в одну строку:
	// hello, worldhello, world

	fmt.Println("hello, world")
	fmt.Print("hello, world")
}
