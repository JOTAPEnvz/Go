package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)
type Users struct{
	Users []User `json:"usuarios"`
}

type User struct{
	Nome string `json:"Nome"`
	Tipo string `json:"Tipo"`
	Idade int `json:"Idade"`
	Social Social `json:"Social"`
}

type Social struct{
	Facebook string `json:"Facebook"`
	Linkedin string `json:"LinkedIn"`
}

func main(){
	jsonFile, err := os.Open("usuarios.json")
	if err != nil{
		fmt.Println("erro abrindo arquivo:", err)
		return
	}
	fmt.Println("Arquvivo aberto!")
	defer jsonFile.Close()

	byteValue, _:= ioutil.ReadAll(jsonFile)

	var usuarios Users

	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.Users);i++{
		fmt.Println("Usuario Tipo: " + usuarios.Users[i].Tipo)
		fmt.Println("Usuario Idade: " + strconv.Itoa(usuarios.Users[i].Idade) )
		fmt.Println("Usuario Nome: " + usuarios.Users[i].Nome)
	}

}