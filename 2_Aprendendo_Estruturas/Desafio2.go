package main

import "fmt"

func main(){

	for x := 1; x <= 100; x++{
		if x %3 == 0 && x %5 == 0{
			fmt.Println("-PinPan", x)
		}else{
			if x %3 == 0{
				fmt.Println("-Pin", x)
			}else{
				if x %5 == 0{
					fmt.Println("-Pan", x)
				}
			}
		}
	}
}