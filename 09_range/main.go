package main

import "fmt"


func main() {
	ids := []int{11,23,44,55,32,64}

	//loop through ids
	for _, id := range ids {
		fmt.Printf("ID: %d\n",  id)
	}

	// add all together

	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Ron":"ron@gmail.com", "Milo":"milo@gmail.com","Mia":"mia@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}