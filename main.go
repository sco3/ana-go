package main

import (
	"ana-go/v2/ana"
	"fmt"
)

func main() {
	given := []string{"asdf", "jkl"}

	pals := ana.Get_anagrams(given)

	fmt.Printf("hi %v\n", pals)

}
