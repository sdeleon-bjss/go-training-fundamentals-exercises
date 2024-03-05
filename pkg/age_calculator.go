package pkg

import (
	"github.com/bearbin/go-age"
	"time"
)

func CalculateAge(dob time.Time) int {
	ageInYears := age.Age(dob)
	return ageInYears
}
