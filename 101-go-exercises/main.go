package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}


func main() {
	//0 create variable doingGoRightNow
	var doingGoRightNow bool = true
	fmt.Println(doingGoRightNow)

	//  Exercise 1
	// On the line below, create a variable named onMarsRightNow and assign it the boolean value of false
	var onMarseRightNow = false;
	fmt.Println(onMarseRightNow)

	// Exercise 2
	// Create a variable named fruits and assign it an array of strings containing the following fruits.
	// mango, banana, guava, kiwi, and strawberry.
	fruitSlice:= []string{"mango", "banana", "guava", "strawberry"}
	fmt.Println(fruitSlice)

	//  Exercise 3
	//  Create a variable named vegetables and assign it an array of fruits containing the following vegetable names as strings:
	//  eggplant, broccoli, carrot, cauliflower, and zucchini
	vegetables:= []string{"eggplant", "broccoli", "carrot", "zucchini"}
	fmt.Println(vegetables)


	// Exercise 4
	// Create a variable named numbers and assign it an array of numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	// ==================================
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	// Exercise 5
	// Add the string "tomato" to the end of the fruits array.
	fruitSlice = append(fruitSlice, "tomato")
	fmt.Println(fruitSlice)

	// Exercise 6
	// add the string "tomato" onto the end of the vegetables array.
	vegetables = append(vegetables, "tomato")
	fmt.Println(vegetables)

	// Exercise 7
	// Given the array of numbers defined below, reverse the array of numbers that you created above.
	var someNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, j :=0, len(someNumbers)-1; i < j; i, j = i+1, j-1 {
		someNumbers[i], someNumbers[j] = someNumbers[j], someNumbers[i]
	}

	fmt.Println(someNumbers)

	// Exercise 8
	// Sort the vegetables in alphabetical order.
	sort.Strings(vegetables)

	fmt.Println(vegetables)

	// Exercise 9
	// Write the code necessary to sort the fruits in reverse alphabetical order

	sort.Strings(fruitSlice)
	fmt.Println(fruitSlice)
	for i, j :=0, len(fruitSlice)-1; i < j; i, j = i+1, j-1 {
		fruitSlice[i], fruitSlice[j] = fruitSlice[j], fruitSlice[i]
	}
	fmt.Println(fruitSlice)

	// Exercise 10
	// Write the code necessary to produce a single array that holds all fruits then all vegetables in the order as they were sorted above.
	// Assign the result to a variable named fruitsAndVeggies.
	var fruitsAndVeggies = append(fruitSlice, vegetables...)
	fmt.Println(fruitsAndVeggies)

	// Exercise 11
	// Write a function definition for a function named addOne that takes in a number and returns that number plus one
	fmt.Println(addOne(3))

	// Exercise 12
	// Write a function definition named isPositive that takes in a number and returns true or false if that number is positive.
	fmt.Println(isPositive(-1))

	// Exercise 13
	// Write a function definition named isNegative that takes in a number and returns true or False if that number is negative.
	fmt.Println(isNegative(-1))

	// Exercise 14
	// Write a function definition named isOdd that takes in a number and returns true or false if that number is odd.
	fmt.Println(isOdd(3))

	// Exercise 15
	// Write a function definition named isEven that takes in a number and returns true or false if that number is even.
	fmt.Println(isEven(1))

	// Exercise 16
	// Write a function definition named identity that takes in any input and returns that input.
	fmt.Println(identity(params{foo: "hello"}))
	fmt.Println(identity(params{goal: true}))

	//Exercise 17  Write a function definition named isPositiveOdd that takes in a number and returns true or false if the value is both greater than zero and odd
	fmt.Println(isPositiveOdd(21))

	//Exercise 18 Write a function definition named isPositiveEven that takes in a number and returns true or false if the value is both greater than zero and even
	fmt.Println(isPositiveEven(32))

	// Exercise 19 Write a function definition named isNegativeOdd that takes in a number and returns true or false if the value is both less than zero and odd.
	fmt.Println(isNegativeOdd(-33))

	// Exercise 20
	// Write a function definition named isNegativeEven that takes in a number and returns true or false if the value is both less than zero and even.
	fmt.Println(isNegativeEven(-12))


}

type params struct {
	foo string
	bar int
	baz float32
	goal bool
}

//Ex 16 function
func identity(params params) params {
	return params
}

//Ex 17 func
func isPositiveOdd(num int) bool {
	return (num % 2 != 0) && (num > 0)
}

//ex 18
func isPositiveEven (num int) bool {
	return (num % 2 == 0) && ( num > 0 )
}
//ex 19
func isNegativeOdd (num int) bool {
	return (num % 2 != 0) && (num < 0)
}

//ex 20
func isNegativeEven (num int) bool {
	return (num % 2 == 0) && (num < 0)
}

func isEven(num int) bool {
	return num % 2 == 0
}

func isOdd(num int) bool {
	return num % 2 != 0
}

func isNegative(num int) bool {
	return num < 0
}

func isPositive(num int) bool {
	return num > 0
}

func addOne(num int) int {
	return num + 1
}
