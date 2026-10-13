package main

import (
	"ana-go/v2/ana"
	"fmt"
)

func main() {
	given := []string{"asdf", "jkl", "fdsa", "abc", "cab"}

	fmt.Printf("Given %v\n", given)

	anas := ana.GetAnagrams(given)

	fmt.Printf("Result %v\n", anas)

}
