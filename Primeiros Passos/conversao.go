package main

import "fmt"

var x int = 10
var y float64 = 20.5

func main() {

	var z = float64(x) + y
	fmt.Printf("%v(%T)",z, z)
	// fmt.Println("A soma de", x, "e", y, "é igual a", z)

}