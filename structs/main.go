package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	/*
		hokage := person{
			firstName: "Naruto",
			lastName:  "Uzumaki",
			contact: contactInfo{
				email: "seventh@leaf.com",
				zip:   11111,
			},
		}

		fmt.Printf(hokage)
	*/

	var babaYaga person
	babaYaga.firstName = "John"
	babaYaga.lastName = "Wick"
	babaYaga.contact = contactInfo{email: "jwick@continental.com", zip: 43823}

	babaYagaPtr := &babaYaga
	//babaYaga.updateNamePtr("Swordsman")
	babaYagaPtr.updateNameValue("Gunslinger")
	babaYaga.print()
	babaYaga.updateNamePtr("Ballerina")
	babaYaga.print()
	//babaYagaPtr.updateNamePtr("Ballerina")
	//babaYaga.print()
}

/* passed by value
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
*/

func (pointerToPerson *person) updateNamePtr(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) updateNameValue(newFirstName string) {
	p.firstName = newFirstName
}
func (p person) print() {
	fmt.Println("")
	fmt.Printf("%+v", p)
}
