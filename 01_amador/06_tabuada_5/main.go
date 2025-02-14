package main

import "fmt"

//Imprima a tabuada do 5 (de 5x1 até 5x10) usando um loop.
func main() {
	for i:=1; i <=10; i+= 1 {
		fmt.Printf("%d x 5 = %d",i, i * 5)
		fmt.Println("")

	}
}