package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	a = 53.2
	b = 456.6
	x := 2.5
	y := 3.3
	fmt.Printf("Value of a = %v", a)
	fmt.Printf("Value of b = %v", b)
	fmt.Printf("Value of x = %v ", x)
	fmt.Printf("Value of y= %v ", y)

	sum := a + b + x + y
	fmt.Printf("Value of Sum = %v", sum)
}
