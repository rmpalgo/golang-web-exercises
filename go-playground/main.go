package main

import (
	"fmt"
	"strconv"
)

//Package visibility
//scope visibility
//Global visibility with capital Letter
var Global int = 32

//refresher exercise
func main() {

	var myFavoriteNumber int
	myFavoriteNumber = 4
	fmt.Println(myFavoriteNumber)

	var myString string = "hello"
	fmt.Println(myString)

	var myNumber float32
	myNumber = 3.14
	fmt.Println(myNumber)

	var x int = 5
	x++
	fmt.Println(x)

	var num int
	num = 32
	fmt.Printf("%v, %T\n", num, num)

	var numFloat float32
	numFloat = float32(x)
	fmt.Printf("%v, %T\n", numFloat, numFloat)

	var t int = 4
	fmt.Printf("%v, %T\n", t, t)

	var tString string
	tString = strconv.Itoa(t)
	fmt.Printf("%v, %T\n", tString, tString)

	//boolean
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

}
