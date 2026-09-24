package main

import "fmt"

func main() {
    var a, b int

    fmt.Scan(&a, &b)

    a, b = b, a

    fmt.Printf("A = %d\n", a)
    fmt.Printf("B = %d\n", b)
}