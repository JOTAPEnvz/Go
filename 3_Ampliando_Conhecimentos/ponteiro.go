package main

import "fmt"

func inicial(xPrt *int){
	*xPrt = 0
}

func main(){
	x := 5
	inicial(&x)
	fmt.Println(x)
}