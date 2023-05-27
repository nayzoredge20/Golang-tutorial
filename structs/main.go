package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // instead we can also just write contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000},
	}
	//jimPointer := &jim
	updateName(&jim, "Jimmy")
	jim.print()
}

func updateName(PointerToPerson *person, newFirstName string) {
	(*PointerToPerson).firstName = newFirstName

}
func (p person) print() {
	fmt.Printf("%+v", p)
}
