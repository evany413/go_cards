package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"} // order matters
	alex := person{firstName: "Alex", lastName: "Anderson"}
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// fmt.Println(alex)
	// fmt.Printf("%+v", jim)

	alex.print()
	jim.print()

	// reference to the memory address of alex
	alexPointer := &alex
	alexPointer.updateNamePointer("Alexa")
	alex.print()

	jim.updateNamePointer("Jimmy")
	jim.print()

}

// send a copy of person
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// send a pointer to person
func (p *person) updateNamePointer(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Println("")
	fmt.Printf("%+v", p)
}
