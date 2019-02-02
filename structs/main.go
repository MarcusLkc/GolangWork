package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "billy",
		contact: contactInfo{
			email:   "Marcus@gmail.com",
			zipCode: 13377,
		},
	}
	jim.updateName("jimmy")
	jim.print()

}
