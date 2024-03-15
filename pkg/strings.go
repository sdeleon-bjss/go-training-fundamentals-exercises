package pkg

import (
	"errors"
	"fmt"
	"strings"
)

func ConcatStrings(strings ...string) string {
	// this could be improved to use strings.Builder
	var result string

	for _, str := range strings {
		result += str
	}

	return result
}

type Name struct {
	First  string
	Middle string
	Last   string
}

func (n *Name) SetFullName(fullName string) error {
	names := strings.Split(fullName, " ")

	if len(names) < 3 {
		return errors.New("invalid full name")
	}

	n.First = names[0]
	n.Middle = names[1]
	n.Last = names[2]

	return nil
}

func (n *Name) SliceName() []string {
	return []string{n.First, n.Middle, n.Last}
}

func (n *Name) GetFullName() string {
	return n.First + " " + n.Middle + " " + n.Last
}

func (n *Name) Print(slicedName []string) {
	fmt.Printf("full-name : %s, middle-name : %s, and surname : %s ", n.GetFullName(), slicedName[1], slicedName[2])
}

func GetFullName(first, middle, last string) string {
	return fmt.Sprintf("%s %s %s", first, middle, last)
}
