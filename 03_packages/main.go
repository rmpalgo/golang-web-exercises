package main

import (
	"fmt"
	"golang-web-exercises/03_packages/strutil"
	"math"
)


func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}