// Package twofer is for exercism
package twofer

import "fmt"

// ShareWith will return a string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
