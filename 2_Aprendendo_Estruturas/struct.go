package main

import "fmt"

type pessoa struct{
	nome string
	idade int
}

func main(){
	fmt.Println(pessoa{"Ana", 54})
	fmt.Println(pessoa{"João", 33})
}