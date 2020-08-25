package main

import "fmt"

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


}
