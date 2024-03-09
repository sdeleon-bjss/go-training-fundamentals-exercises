package pkg

type Person interface {
	GetFullName() string
	GetDateOfBirth() string
	GetAge() int
}

type Pupil struct {
	FullName    Name
	DateOfBirth string
	Age         int
}

func (p Pupil) GetFullName() string {
	return p.FullName.GetFullName()
}

func (p Pupil) GetDateOfBirth() string {
	return p.DateOfBirth
}

func (p Pupil) GetAge() int {
	return p.Age
}

func NewPupil(firstName, middleName, lastName, dateOfBirth string, age int) Pupil {
	return Pupil{
		FullName: Name{
			First:  firstName,
			Middle: middleName,
			Last:   lastName,
		},
		DateOfBirth: dateOfBirth,
		Age:         age,
	}
}
