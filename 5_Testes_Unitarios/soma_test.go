package main

import "testing"

func TestSoma(t *testing.T){
	teste := soma(4, 5, 6)
	resultado := 15
    
	if teste != resultado{
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}

func TestMult(t *testing.T){
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado{
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}