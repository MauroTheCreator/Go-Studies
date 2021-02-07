package main

import "fmt"

func main() {
	var input float32

	fmt.Scanf("%f", &input)
	var result = input / 3.281

	fmt.Printf("%f feets == %f metters", input, result)
}