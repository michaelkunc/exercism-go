package luhn

import (
	"strconv"
	"strings"
)

// Valid checks the input according to the luhn formula
func Valid(input string) bool {
	sum := 0
	splitInput := strings.Split(input, "")
	for index, char := range splitInput {
		integer, _ := strconv.Atoi(char)
		if (index+1)%2 == 0 {
			sum += integer * 2
		} else {
			sum += integer
		}
	}
	if (sum % 10) == 0 {
		return true
	}
	return false
}
