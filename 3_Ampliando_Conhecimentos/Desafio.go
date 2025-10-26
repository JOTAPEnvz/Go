package main

import (
	"fmt"
	"time"
)	

func ping(c chan string){
	for i := 0; i == 0; i++{
	    c <- "ping"
		fmt.Println("ping")
	}
	for{
		msg:= <-c
		if msg == "pong"{
			time.Sleep(time.Second * 1)
			fmt.Println("ping")
			c <- "ping"
		}
	}
}

func pong(c chan string){
	for {
		msg:= <-c
		if msg == "ping"{
		   time.Sleep(time.Second * 1)
		   fmt.Println("pong")
		   c <- "pong"
	    }
	}
}

func main(){
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)

	var entrada string
	fmt.Scanln(&entrada)
}