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

func main() {
	contact := contactInfo{email: "jim@jimmy.com", zipCode: 190339}
	jim := person{firstName: "jim", lastName: "carry", contact: contact}
	jim.updateName("jimCarry")
	jim.print()

	mySlice := []string{"hello", "world"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
