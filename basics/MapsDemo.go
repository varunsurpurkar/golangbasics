package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float32{
		"TESLA": 234.56,
		"BOFA":  45.56,
		"CFG":   41.32,
		"APPL":  451.23,
	}

	fmt.Println(stocks)

	fmt.Println(stocks["APPL"])

	stocks["WFS"] = 34.34

	// key value looping
	for key, value := range stocks {
		fmt.Printf("Key = %s Value = %f \n", key, value)
	}

	// key looping range

	for key := range stocks {
		fmt.Printf("key = %s \n", key)

	}

	// key looping range

	for _, value := range stocks {
		fmt.Printf("value = %f \n", value)

	}

}
