package main

import "testing"

func TestSoma2(t *testing.T){
	teste := somar(4, 5, 6)
	resultado := 15
    
	if teste != resultado{
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}

func TestSubtrai(t *testing.T){
	teste := subtrair(10, 5)
	resultado := 5

	if teste != resultado{
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMult2(t *testing.T){
	teste := multiplicar(10, 10)
	resultado := 100

	if teste != resultado{
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}

func TestDividir(t *testing.T){
	teste := dividir(10, 2)
	resultado := 5.0
	
	if teste != resultado{
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

