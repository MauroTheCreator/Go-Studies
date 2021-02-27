package main

import "fmt

func halves(n int) (res int, isEven bool) {
	if n % 2 == 0 {
		res, isEven = n / 2, true
	} else { 
		res, isEven = 0, false
	}
	return
}

func main() {
	res, isEven := halves(1)
	fmt.Println(res, isEven)
}
