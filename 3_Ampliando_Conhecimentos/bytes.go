package main

import (
	"bytes"
	"fmt"
)

func main(){
	fmt.Printf("%s", bytes.Title([]byte("Testando pacote bytes \n")))
}