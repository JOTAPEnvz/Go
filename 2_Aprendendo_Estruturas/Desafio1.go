package main

import "fmt"

func main(){

	arr := []int{}
	
	for x:= 1; x <= 100; x++{
		if x %3 == 0{
			arr = append(arr, x)
		}
		if x >= 100{
			fmt.Println("Os números divisiveis por 3 de 1 a 100 são:", arr) 
		}
	}

}