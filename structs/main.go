package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	alex.firstName = "Vinh"
	alex.lastName = "Bui"

	fmt.Println(alex)
	fmt.Printf("+%v", alex)

}
