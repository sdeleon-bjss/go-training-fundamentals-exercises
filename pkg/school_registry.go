package pkg

type Pupil struct {
	FullName    Name
	DateOfBirth string
	Age         int
}

func (p Pupil) GetFullName() string {
	return p.FullName.GetFullName()
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
