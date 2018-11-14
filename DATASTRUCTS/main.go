package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p := person{
		firstName: "Jim",
		lastName:  "Smith",
	}

	p.print()

	ptrPerson := &p
	ptrPerson.updateFirstName("Jimmy")

	p.print()

	ptrPerson.updateFirstName("Fred")

	p.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}
