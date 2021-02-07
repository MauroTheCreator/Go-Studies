package main

import "fmt"

func main() {
	var input float64 // entrada do usuário
	
	fmt.Scanf("%f", &input) // usuário coloca os dados no programa
	var result float64 = (input - 32) * 5 / 9

	fmt.Printf("%f fahr == %f celsius", input, result)

}
