package main

import (
	"fmt"
	"prose"
)

func main() {
	phrases := []string{"apple"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
