package main

import "fmt"

func main(){
    x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])

	inteiro()
	elemento()
}

func inteiro(){
	x := make(map[int]int)
	x[1] = 20
	x[2] = 30
	fmt.Println(x[1], x[2])
}

func elemento(){
    x := make(map[string]string)
	x["Fe"] = "Ferro" 
	x["U"] = "Urânio"
	x["Au"] = "Ouro"
    
	fmt.Println(x["Fe"], x["U"], x["Au"])
}