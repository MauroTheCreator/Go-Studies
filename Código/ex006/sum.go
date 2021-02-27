package main

import "fmt"

func sum(slice ...int) (res int) {
	for _, number := range slice {
		res += number
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
