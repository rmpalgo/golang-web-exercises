package main

import "fmt"

//Package visibility
//scope visibility
//Global visibility with capital Letter
var Global int = 32

//refresher exercise
func main() {

	//var myFavoriteNumber int
	//myFavoriteNumber = 4
	//fmt.Println(myFavoriteNumber)
	//
	//var myString string = "hello"
	//fmt.Println(myString)
	//
	//var myNumber float32
	//myNumber = 3.14
	//fmt.Println(myNumber)
	//
	//var x int = 5
	//x++
	//fmt.Println(x)
	//
	//var num int
	//num = 32
	//fmt.Printf("%v, %T\n", num, num)
	//
	//var numFloat float32
	//numFloat = float32(x)
	//fmt.Printf("%v, %T\n", numFloat, numFloat)
	//
	//var t int = 4
	//fmt.Printf("%v, %T\n", t, t)
	//
	//var tString string
	//tString = strconv.Itoa(t)
	//fmt.Printf("%v, %T\n", tString, tString)
	//
	////boolean
	//n := 1 == 1
	//m := 1 == 2
	//fmt.Printf("%v, %T\n", n, n)
	//fmt.Printf("%v, %T\n", m, m)

	a := 10 // 1010
	b := 3 // 0011

	fmt.Println(a & b) // 0010 & (bitwise AND): Takes two numbers as operands and does AND on every bit of two numbers. The result of AND is 1 only if both bits are 1.
	fmt.Println(a | b) // 1011 Takes two numbers as operands and does OR on every bit of two numbers. The result of OR is 1 any of the two bits is 1.
	fmt.Println(a ^ b) // 1001 Takes two numbers as operands and does XOR on every bit of two numbers. The result of XOR is 1 if the two bits are different.
	fmt.Println(a &^ b) // 0100 This is a bit clear operator.

	c := 8
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2 ^ 0 or 1



}
