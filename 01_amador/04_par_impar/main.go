package main

import "fmt"

//Peça um número ao usuário e diga se ele é par ou ímpar.
func main() {
	var number int
	fmt.Print("Digite um número inteiro e eu te falo se é par ou impar: ")
	fmt.Scan(&number)

	if number % 2 == 0 {
		fmt.Printf("%d é um número par", number)
	} else {
		fmt.Printf("%d é um número impar", number)
	}
}