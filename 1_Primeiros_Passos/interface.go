package main

import "fmt"

type Stringer interface {
	String() string
}

type Autor struct {
	Nome string
}
func (a Autor) String() string {
	return fmt.Sprintf("Eu %s estou aprendendo Go", a.Nome)
}

func main() {
	a := Autor{
		Nome: "João Paulo",
	}
	fmt.Println(a.String())
}