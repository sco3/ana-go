package ana

import "fmt"

func GetAnagrams(given []string) [][]string {
	result := [][]string{}
	ana_key_map := BuildMap(given)

	fmt.Printf("Map: %v\n", ana_key_map)

	for i, words := range ana_key_map {
		fmt.Printf("%v %v\n", i, words)

		if len(words) > 1 {
			result = append(result, words)
		}

	}
	return result
}
