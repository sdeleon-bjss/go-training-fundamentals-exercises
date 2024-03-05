package pkg

type Name struct {
	First  string
	Middle string
	Last   string
}

func (n Name) GetFullName() string {
	return n.First + " " + n.Middle + " " + n.Last
}
