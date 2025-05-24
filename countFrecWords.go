package tripofgo

import (
	"strings"
)

func CountsFrecWords(s string) map[string]int {
	// convert string to []string
	words := strings.Fields(s)

	// make map for store words
	mp := make(map[string]int)

	for _, word := range words {
		mp[word]++
	}

	return mp
}
