package main

import "fmt"

type contractInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contractInfo
}

func main() {
	jhon := person{
		firstName: "Jhon",
		lastName:  "Doe",
		contractInfo: contractInfo{
			email:   "jhondoe@email.com",
			zipCode: 9321,
		},
	}
	jhon.updateName("Jim")
	jhon.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
