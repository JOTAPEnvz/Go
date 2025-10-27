package main

import (
	"fmt"
)

func main(){
	x := somar(4, 5 ,6)
	y := multiplicar(10, 10)
	z := dividir(10, 3)
	w := subtrair(10, 5, 2)
	fmt.Println(x, y, z, w)
}

func somar(i ...int)int{
	total := 0

	for _, v := range i{
		total += v
	}


	return total
}

func subtrair(i ...int) int {
	result := i[0]
	for _, v := range i[1:] {
		result = result - v
	}
	return result
}

func multiplicar(i ...int)int{
	total := 1
	for _, v := range i{
		total = total * v
	}
	return total
}

func dividir(i ...int) float64 {
	if len(i) == 0 {
		return 0
	}
	result := float64(i[0])
	for _, v := range i[1:] {
		if v == 0 {
			return 0
		}
		result = result / float64(v)
	}
	return result
}