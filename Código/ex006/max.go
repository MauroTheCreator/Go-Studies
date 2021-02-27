package main

import "fmt"

func max(values ...int) (mx int) {
	mx = values[0]
	for _, v := range values {
		if v > mx { mx = v }
	}
	return
}

func main() {
	fmt.Println(max(1, 2, 3, 6, 7, 3, 1, 0))
}