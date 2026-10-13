package main

import (
	"ana-go/v2/ana"
	"fmt"
)

func main() {
	given := []string{"asdf", "jkl"}

	anas := ana.Get_anagrams(given)

	fmt.Printf("hi %v\n", anas)

}
