package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
	}

	jim.print()

	ptrJim := &jim

	ptrJim.updateFirstName("Jimmy")

	jim.print()

	ptrJim.updateFirstName("Jimmed")

	jim.print()

	(*(&jim)).updateFirstName("Jimmer")

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}
