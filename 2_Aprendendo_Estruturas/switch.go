package main

import "fmt"

func main() {
	for i := 0; i <=10; i++ {
		switch i %2{
		case 0: fmt.Println(i, "é Par")
		case 1: fmt.Println(i, "é Impar")
		}
	}
}