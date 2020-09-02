package main

import "fmt"

//arrays, slices, maps, have to have same data types whereas struct do not



//Structs Review
type Doctor struct {
	number int
	actorName string
	companions []string
	episodes []string
}


//What are they?
//Creating
//CreatingNaming conventions
//Embedding
//Tags


func main() {
	//best method is to place field names
	aDoctor := Doctor{
		number : 3,
		actorName: "Jon Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
fmt.Println(aDoctor)
}