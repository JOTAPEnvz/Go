package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {
	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9

	fmt.Println(tempF, "graus Fahrenheit é igual a", tempC, "graus Celsius")


}
