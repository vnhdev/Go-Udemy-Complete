package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	vinh := person{
		firstName: "Vinh",
		lastName:  "Bui",
		contactInfo: contactInfo{
			email:   "wingsolobot@gmail.com",
			zipCode: 700000},
	}
	vinh.updateName("Phong")
	vinh.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
