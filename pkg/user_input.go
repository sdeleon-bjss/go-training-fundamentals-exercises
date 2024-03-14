package pkg

import "fmt"

// GetFullName takes in three strings and returns a full name
func GetFullName(first, middle, last string) string {
	return fmt.Sprintf("%s %s %s", first, middle, last)
}
