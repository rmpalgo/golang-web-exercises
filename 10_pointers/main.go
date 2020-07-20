package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)
	// pointer => b has *int type two different data types
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change val wit pointer
	*b = 10
	fmt.Println(a)
	// go is passed by value

}