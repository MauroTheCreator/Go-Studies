package main

import "fmt"

func makeOddGen() func() uint {
	i := uint(1)
	return func() (res uint) {
		res = i
		i += 2
		return
	}
}


func main() {
	nextOdd := makeOddGen()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}