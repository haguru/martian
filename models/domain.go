package models

type Domain struct {
	Name         string
	Members      []Person
	OncallPerson Person
}
