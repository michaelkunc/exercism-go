// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimRemark := strings.TrimSpace(remark)
	isUpper := strings.ToUpper(trimRemark) == trimRemark
	isQuestion := strings.HasSuffix(trimRemark, "?")
	isExclamation := strings.HasSuffix(trimRemark, "!")
	anyLettersRe, _ := regexp.Compile("[a-zA-Z]")
	anyLetters := anyLettersRe.MatchString(trimRemark)
	isSilence := len(trimRemark) == 0
	switch {
	case isUpper && isQuestion && anyLetters:
		return "Calm down, I know what I'm doing!"
	case (isUpper && isExclamation) || (isUpper && anyLetters):
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case isSilence:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
