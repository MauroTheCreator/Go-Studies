package main

import "fmt"
import "time"

func f() {
    fmt.Println("function")
}

func main() {
    f()

    for i := 0; i < 11; i++ {
        go f()
    }

    for i := 0; i < 11; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}
