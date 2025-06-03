package main

func main() {
	println(len("Hello World")) // Conta o o número de caracteres na string
	println("Hello World"[0]) // Mostra o caractere na posição 0(como go lê em binario será um número)  
	println(string("Hello World"[0])) // Pega o caractere e não somente sua representação binaria
	println("Hello" + "World") // Concatena duas strings
}