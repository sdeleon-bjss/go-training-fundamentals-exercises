package pkg

import (
	"errors"
	"github.com/bearbin/go-age"
	"time"
)

type Student interface {
	GetFullName() string
	GetDateOfBirth() string
	GetAge() int
}

type Pupil struct {
	FullName Name
	DOB      time.Time
}

func (p Pupil) GetFullName() string {
	return p.FullName.GetFullName()
}

func (p Pupil) GetDateOfBirth() string {
	return p.DOB.Format("01/02/2006")
}

func (p Pupil) GetAge() int {
	return age.Age(p.DOB)
}

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
