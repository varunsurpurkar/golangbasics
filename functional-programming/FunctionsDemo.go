package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func div(x int, y int) (int, int, error) {
	if x == 0 {
		return 0, 0, fmt.Errorf("Numerator must not be zero %d", x)
	}
	return x / y, x % y, nil
}

func cleanUp(value string) {
	fmt.Printf("CLEANING = %s", value)
}

func main() {
	defer cleanUp("ALL RESOURCES \n")

	val := add(1, 2)
	fmt.Printf("Sum = %d", val)
	var quotient int
	var remainder int
	var err error
	quotient, remainder, err = div(9, 3)
	if err != nil {
		fmt.Printf("Error Message = %s ", err)
	} else {
		fmt.Printf("\n quotient = %d remainder = %d \n", quotient, remainder)
	}

}
