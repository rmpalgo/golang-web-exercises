package main

import "fmt"

func main() {
	x := 15
	y := 10

	//if else statements
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if

	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("Idk the color")
	}

	//Switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "bue":
		fmt.Println("color is red")
	default:
		fmt.Println("idk the color")
	}

}