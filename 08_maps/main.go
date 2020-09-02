package main

import "fmt"

func main() {
	// Maps
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	//m := map[[3]string]string{}
	//return of map is not guarantee the order
	statePopulations["Georgia"] = 10310371
	//the "ok" will let you know if the key is indeed a key in the map
	//since non key value will still return 0 as default
	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))

	//when referencing map, will affect original data
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

}