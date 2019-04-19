// Package scrabble for scoring words in scrabble
package scrabble

import (
	"strings"
)

// Score returns the scrabble score for a given word
func Score(n string) (total int) {
	total = 0
	var letterMapping = map[string]int{
		"a": 1,
		"e": 1,
		"i": 1,
		"o": 1,
		"u": 1,
		"l": 1,
		"n": 1,
		"r": 1,
		"s": 1,
		"t": 1,
		"d": 2,
		"g": 2,
		"b": 3,
		"c": 3,
		"m": 3,
		"p": 3,
		"f": 4,
		"h": 4,
		"v": 4,
		"w": 4,
		"y": 4,
		"k": 5,
		"z": 10,
	}
	lowerCase := strings.ToLower(n)
	for _, c := range lowerCase {
		letterScore := letterMapping[string(c)]
		total += letterScore

	}
	return total
}
