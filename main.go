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
	var color1 map[string]string
	fmt.Println(color1)
	colors2 := make(map[string]string)
	fmt.Println(colors2)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	delete(colors, "red")
	printMap(colors)

	f := fibonacci()
	for i := 0; i < 10; i++ {
		if i == 0 || i == 1 {
			fmt.Println(i)
		} else {
			fmt.Println(f())
		}
	}
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
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

// function closure
// print fibonnaci

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}
