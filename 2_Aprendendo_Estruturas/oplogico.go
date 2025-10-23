package main

import "fmt"

func main()	{
	X:= 6
	Y:= 7

	if X %2 == 0 || X %3 == 0{
		fmt.Println("X é válido")
	}

	if Y %2 == 0 || Y %3 == 0{
		fmt.Println("Y é válido")
	}else{
		fmt.Println("Y é inválido")
	}

}