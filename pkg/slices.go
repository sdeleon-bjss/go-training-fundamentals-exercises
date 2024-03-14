package pkg

import (
	"errors"
	"fmt"
	"strings"
)

// Name is a struct that holds the first, middle, and last name
type Name struct {
	First  string
	Middle string
	Last   string
}

// SetFullName method sets the full name of the Name struct
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

// GetFullName method returns the full name of the Name struct
func (n *Name) GetFullName() string {
	return n.First + " " + n.Middle + " " + n.Last
}

// Print method prints out the contents in format for exercise
func (n *Name) Print() {
	fmt.Printf("full-name : %s, middle-name : %s, and surname : %s ", n.GetFullName(), n.Middle, n.Last)
}
