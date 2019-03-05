// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

const letters = "abcdefghijklmnopqrstuvwxyz"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	isUpper := strings.ToUpper(remark) == remark
	isQuestion := strings.HasSuffix(remark, "?")
	isExclamation := strings.HasSuffix(remark, "!")
	switch {
	case isUpper && isQuestion && onlyLetters(remark):
		return "Calm down, I know what I'm doing!"
	case (isUpper && isExclamation) || (isUpper && onlyLetters(remark)):
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

func onlyLetters(phrase string) bool {
	for _, char := range phrase {
		if !strings.Contains(letters, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
