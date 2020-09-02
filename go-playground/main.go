package main

import "fmt"

//infer the pattern
const (
	a = iota
	b
	c
	)

const (
	a2 = iota
)

const (
	errorSpecialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)
func main() {
	//const a = 42
	//var b int16 = 27
	//fmt.Printf("%v, %T\n", a + b, a + b)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)

	var specialistType int
	fmt.Println(specialistType)
	fmt.Printf("%v\n", specialistType == catSpecialist)

}
