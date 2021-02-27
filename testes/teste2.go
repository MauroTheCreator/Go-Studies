package main

import "fmt"

func count(until int, msg string) {
    for i := 0; i <= until; i++ {
        fmt.Println(i)
    }
    fmt.Println(msg)
}

func main() {
    count(10, "finished")
    
    go count(10, "finished")

    fmt.Println("going")
}
