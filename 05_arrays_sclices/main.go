package main

import "fmt"

func main() {

	grades := [...]int{97, 85, 93, 94}
	var students [3]string
	fmt.Printf("Grades: %v\n", grades)
	students[0] = "Milo"
	students[1] = "Mia"
	students[2] = "Lucy"
	fmt.Printf("Grades: %v\n", students)
	fmt.Printf("Grades size: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	//arrays are values, reassigning leads to a copy of the array, imagine hundreds of thousands element being copied
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// using & as pointers will point to the same data as a has
	c := &a
	c[1] = 0
	fmt.Println(a)
	fmt.Println(c)
}
