// Package hamming returns the hamming distance for two strings
package hamming

import "errors"

// Distance calculates the hamming distance
func Distance(a, b string) (int, error) {
	hamming := 0
	if len(a) != len(b) {
		return -1, errors.New("strings must be same length")
	}
	for i := range a {
		if a[i] != b[i] {
			hamming = hamming + 1
		}
	}
	return hamming, nil
}
