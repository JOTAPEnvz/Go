package main

import(
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "API")
}

func main(){
	http.HandleFunc("/home", home)

	http.ListenAndServe(":8080", nil)
}