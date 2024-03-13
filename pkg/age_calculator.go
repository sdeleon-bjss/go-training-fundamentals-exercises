package pkg

import (
	"github.com/bearbin/go-age"
	"time"
)

type Person struct {
	DOB time.Time
}

func (p *Person) Age() int {
	return age.Age(p.DOB)
}
