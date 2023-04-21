package main

// import "errors"

type Person struct {
	age int
}

func NewPerson(age int) (*Person) {

	if age < 1 {
		return nil
	}

	return &Person{
		age: age,
	}
}

func (p *Person) older(other *Person) bool {
	return p.age > other.age
}