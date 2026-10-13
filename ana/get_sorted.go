package ana

import "slices"

func GetSorted(s string) string {
	r_word := []rune(s)
	slices.Sort(r_word)
	return string(r_word)
}
