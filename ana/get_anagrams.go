package ana

import "fmt"

func Get_anagrams(given []string) []string {
	ana_key_map := Build_map(given)

	fmt.Printf("Map: %v\n", ana_key_map)

	r := given //[]string{}
	return r
}
