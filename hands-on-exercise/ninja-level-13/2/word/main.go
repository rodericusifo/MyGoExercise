// package word helps you to do operation related to a bunch of words
package word

import (
	"strings"
)

// UseCount will count every word in your sentence and return it to a map
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count will calculate how many words in your sentence
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
