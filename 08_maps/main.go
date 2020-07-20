package main

import "fmt"

func main() {
	////Define map
	//
	//emails := make(map[string]string)
	//
	//// Assign kv
	//emails["Ron"] = "ron@gmail.com"
	//emails["Milo"] = "milo@gmail.com"
	//emails["Mia"] = "mia@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Ron":"ron@gmail.com", "Milo":"milo@gmail.com","Mia":"mia@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Milo"])

	// Delete from map
	delete(emails, "Ron")
	fmt.Println(emails)
}