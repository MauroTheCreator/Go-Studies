package main

import "fmt"

// versão não recursiva dos números de fibonacci
// func fib(n int) int {
// 	a, b, c, i := 0, 1, 0, 0
// 	for i < n {
// 		c = a + b
// 		a = b
// 		b = c
// 		i++
// 	}
// 	return c
// }

// versão recursiva da função fib(n)
func fib(n uint) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}


func main() {
	fmt.Println(fib(6))
}