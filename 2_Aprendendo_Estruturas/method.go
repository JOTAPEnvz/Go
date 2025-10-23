package main

import "fmt"

type retangulo struct{
	comprimento, altura int
}

func (r *retangulo) area() int{
	return r.comprimento * r.altura
}

func (r *retangulo) perimetro() int{
	return 2*r.comprimento + 2*r.altura
}

func main(){
	r := retangulo{comprimento:10, altura: 20}
	fmt.Printf("Area= %d cm²\n", r.area())
	fmt.Printf("Perimetro= %d cm\n", r.perimetro())
}