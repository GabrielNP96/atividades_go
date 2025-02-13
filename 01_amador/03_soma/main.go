package main

import "fmt"

func main() {
	var firstNumber int
	var secondNumber int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&firstNumber)
	fmt.Print("digite outro número inteiro: ")
	fmt.Scan((&secondNumber))

	var sun int = firstNumber + secondNumber

	fmt.Printf("%d + %d = %d", firstNumber, secondNumber, sun)
}