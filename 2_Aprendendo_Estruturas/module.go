package main

import "fmt"

func main(){
    for x := 10; x > 2;{
	  fmt.Println(x)
	  x = (x/2)
    }
}