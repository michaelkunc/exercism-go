// Package acronym returns 3 letter acronyms
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var letters []string
	wordList := strings.Fields(strings.ReplaceAll(s, "-", " "))
	for _, v := range wordList {
		firstLetter := string(v[0])
		letters = append(letters, strings.ToUpper(firstLetter))
	}
	return strings.Join(letters, "")

}
