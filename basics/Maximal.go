package main

import (
	"fmt"
)

func main() {

	numbers := []int{16, 84, 42, 3434, 15}

	max := numbers[0]

	for _, value := range numbers[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Printf("/n Maximum value = %d ", max)
}
