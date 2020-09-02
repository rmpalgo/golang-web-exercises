package main

import "fmt"

func main() {
/*
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
 */

	/*
	//slice review
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	 */

	/*
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:] //slice of all elements
	b[0] = 0
	c := a[3:] //slice from 4th element to the end
	d := a[:6] //slice first 6 elements
	e := a[3:6] //slice the 4th, 5th, and 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	 */

	/*
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	 */

	/*
	b := []int{2, 3,4, 5}
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, b...)
	fmt.Println(a)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	 */

	//shift, pop
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[1:] //shift
	c := a[:len(a) - 1] //pop
	d := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a)


}
