package main

import (
	"fmt"
)

func main() {

	sentence := "VARUNSURPURKAR"
	lesson := `a quick brown fox jumped on a lazy dog`
	fmt.Println(len(sentence))
	fmt.Printf("sentence = %v\n", sentence[0])
	fmt.Println(sentence[5:])
	fmt.Println(sentence[0:5])
	fmt.Println(sentence[:5])
	fmt.Println(lesson)
}
