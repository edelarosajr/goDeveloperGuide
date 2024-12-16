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

// Go is "Pass by Value"
// a slice is a pointer variable
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

}

/* passed by value
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
*/

/*
# Learnings on Pointers
 1. You do not need to use an asterisk* to refer to a value in a function that had a pointer as a receiver
 2. When it comes to functions, think that the function can either change the original value (ptr receiver)
    or modifies a copy (value receiver). How you call it is irrelevant an Go will dereference where necessary.
 3. In the end what matters is point 2: think from the POV of the function and what needs changing
 4. Printing the value of a pointer for a primitive type (int, string, bool) will print the address;
    for a complex value like struct will print the underlying value as the address would be too complex
 5. A slice is really a struct with (ptr head, capacity, and lenght). When passed to a value function, a copy is made,
    but the copy is still pointing to the same value.
 6. Value types: int, float, string, booo, structs. Use PTRs to change in a function
 7. reference types: slices, maps, channels, pointers, functions. Dont need to worry about PTRs to change in a function
*/
func (pointerToPerson *person) updateNamePtr(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}
func (p person) updateNameValue(newFirstName string) {
	p.firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}
