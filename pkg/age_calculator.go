package pkg

import (
	"github.com/bearbin/go-age"
	"time"
)

// Person struct
//
// DOB: Date of Birth in time.Time format
type Person struct {
	DOB time.Time
}

// Age method returns the age of the person
//
// uses the go-age pkg
func (p *Person) Age() int {
	return age.Age(p.DOB)
}
