package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	even1:= makeEvenGenerator()
	even2:= makeEvenGenerator()

	fmt.Println(even1(), even2(), even1(), even2())
}
