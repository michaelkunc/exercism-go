// Package raindrops is for the raindrops exercise
package raindrops

// Convert does the coversion
func Convert(n int) string {
	var result string
	if n%3 == 0 {
		result += "Pling"
	} else {
		result += string(n)
	}

	return result
}
