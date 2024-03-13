package pkg

import "fmt"

func GetFullName(first, middle, last string) string {
	return fmt.Sprintf("%s %s %s", first, middle, last)
}
