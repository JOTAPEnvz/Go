// Converte o ponto de ebulição da água de Kelvin para Celsius
package main

import "fmt"

const ebulicaoK float64 = 373.2

func main() {
	var tempC = ebulicaoK - 273
	fmt.Println(ebulicaoK, "graus Kelvin é igual a", tempC, "graus Celsius")
	fmt.Printf("%.2f graus Kelvin é igual a %.2f graus Celsius\n", ebulicaoK, tempC)
}