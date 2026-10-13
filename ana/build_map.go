package ana

func BuildMap(given []string) map[string][]string {
	result := make(map[string][]string)
	for _, s := range given {
		key := GetSorted(s)
		result[key] = append(result[key], s)
	}
	return result
}
