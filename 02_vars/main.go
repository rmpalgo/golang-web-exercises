package main

import "fmt"



func main() {
	//MAIN TYPES

	//String
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	// var keyword to create variable
	var age int32 = 37
	var isCool = true
	isCool = false
	//var size float32 = 2.3
	//shorthand declaring variable
	//name := "Ron"
	//email := "ron@gmail.com"
	size := 1.3

	name, email := "Ron", "ron@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)


}