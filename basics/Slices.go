package main

import "fmt"

func main() {

	fruits := []string{"Apple", "Mango", "Banana", "Orange", "WaterMelon"}

	fmt.Println(fruits)

	fmt.Println(fruits[0])

	// starting from index 1
	fmt.Println(fruits[1:])

	//uptil index 4
	fmt.Println(fruits[:4])

	//index 2 to 4

	fmt.Println(fruits[2:4])

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fmt.Println("-------------")

	for i, name := range fruits {
		fmt.Println("*******")
		fmt.Println(i)
		fmt.Println(name)
		fmt.Println("*******")
	}

	for i := range fruits {
		fmt.Println("!!!!!!!!")
		fmt.Println(i)
		fmt.Println("!!!!!!!!!!")
	}

	for _, name := range fruits {
		fmt.Println("%%%%%%%%%%%%")
		fmt.Printf("Name = %s\n", name)
		fmt.Println("%%%%%%%%%%%%")

	}
}
