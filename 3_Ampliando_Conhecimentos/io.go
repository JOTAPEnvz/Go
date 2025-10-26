package main

import(
	"io"
	"log"
	"os"
)

func main(){
	if _, err := io.WriteString(os.Stdout, "Hello World \n"); err != nil {
		log.Fatal(err)
	}
}