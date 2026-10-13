package ana

func Build_map(given []string) map[string][]rune {
	result := make(map[string][]rune, len(given))
	for _, s := range given {
		word := []rune(s)
		result[s] = word

	}

	return result

}
