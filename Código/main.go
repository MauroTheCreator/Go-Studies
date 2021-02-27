package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	a := float64(10)
	square(&a)
	fmt.Println(a)
}
