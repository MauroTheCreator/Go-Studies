package main

import "fmt"

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	a, b := 1, 2
	fmt.Println(a, b)

	swap(&a, &b)
	fmt.Println(a, b)
}