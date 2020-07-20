package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age								  int
}

//greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}


}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Ron", lastName: "Pal", city: "DFW", gender: "m", age: 123 }
	//
	////alternative
	////person1 := Person{"Ron", "Pal", "DFW",  "m", 25}
	//fmt.Println(person1)
	//fmt.Println(person1.firstName)
	//person1.age++
	//fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Hello")
	fmt.Println(person1.greet())

}