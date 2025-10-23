package main

import(
	"fmt" 
	"math"
)

type geometria interface{
	area() float64
	perimetro() float64
}

type quadrado struct{
	lado float64
}

type circulo struct{
	raio float64
}

func (q quadrado) area() float64{
	return q.lado * q.lado
}

func (q quadrado) perimetro() float64{
	return q.lado * 4
}

func (c circulo) area() float64{
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64{
	return 2 * math.Pi * c.raio
}

func calcular(g geometria){
	fmt.Println(g)
	fmt.Printf("Area: %f\n", g.area())
	fmt.Printf("Perimetro: %f\n", g.perimetro())
}

func main(){
	q := quadrado{lado:10}
	c := circulo{raio:5}
	calcular(q)
	calcular(c)
}