// Package raindrops is for the raindrops exercise
package raindrops

import "strconv"

// Convert does the coversion
func Convert(n int) (result string) {
	var dropMapping = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	for num, noise := range dropMapping {
		if n%num == 0 {
			result += noise

		}
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result
}
