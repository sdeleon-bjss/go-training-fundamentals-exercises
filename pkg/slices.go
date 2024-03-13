package pkg

import (
	"errors"
	"strings"
)

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

func (n *Name) GetFullName() string {
	return n.First + " " + n.Middle + " " + n.Last
}
