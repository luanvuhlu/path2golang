package main

import "fmt"

type contact struct {
	address string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contact
}

func newPerson() person {
	return person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contact{
			address: "Hanoi",
			zipCode: 999,
		},
	}
}
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) updateFirstName(jimmy *person, name string) {
	jimmy.firstName = name
	(*p).firstName = name
}
