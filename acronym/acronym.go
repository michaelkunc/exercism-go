// Package acronym returns 3 letter acronyms
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var letters []string
	wordList := strings.Split(s, " ")
	for _, v := range wordList {
		letters = append(letters, string(v[0]))
	}
	return strings.Join(letters, "")

}
