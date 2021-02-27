package main

import "fmt"

func point(variable int) (ptr *int) {
	return &variable
}


func main() {
	a := 10
	b := point(a)
	fmt.Println(*b)
}