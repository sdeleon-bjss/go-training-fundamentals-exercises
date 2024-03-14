package pkg

import (
	"errors"
	"github.com/bearbin/go-age"
	"time"
)

// Student interface
type Student interface {
	GetFullName() string
	GetDateOfBirth() string
	GetAge() int
}

// Pupil struct
type Pupil struct {
	FullName Name
	DOB      time.Time
}

// GetFullName method returns the full name of the pupil
//
// this leverages the GetFullName method from the Name struct
func (p Pupil) GetFullName() string {
	return p.FullName.GetFullName()
}

// GetDateOfBirth method returns the date of birth of the pupil
func (p Pupil) GetDateOfBirth() string {
	return p.DOB.Format("01/02/2006")
}

// GetAge method returns the age of the pupil
//
// uses the go-age package
func (p Pupil) GetAge() int {
	return age.Age(p.DOB)
}

// NewPupil function creates a new pupil
func NewPupil(firstName, middleName, lastName, dateOfBirth string) (Pupil, error) {
	if len(firstName) == 0 || len(middleName) == 0 || len(lastName) == 0 {
		return Pupil{}, errors.New("invalid full name")
	}

	dob, err := time.Parse("01/02/2006", dateOfBirth)
	if err != nil {
		return Pupil{}, err
	}

	pupil := Pupil{FullName: Name{First: firstName, Middle: middleName, Last: lastName}, DOB: dob}

	return pupil, nil
}
