package main

import "fmt"

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

}