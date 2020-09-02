package main

import "fmt"

//arrays, slices, maps, have to have same data types whereas struct do not



//Structs Review
type Doctor struct {
	Number int
	ActorName string
	Companions []string
	Episodes []string
}


//What are they?
//Creating
//CreatingNaming conventions
//Embedding
//Tags


func main() {
	//best method is to place field names
	aDoctor := Doctor{
		Number : 3,
		ActorName: "Jon Pertwee",
		Companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
fmt.Println(aDoctor.Companions[1])
}