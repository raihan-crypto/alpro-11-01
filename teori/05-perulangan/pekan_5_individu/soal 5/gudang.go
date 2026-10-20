package main

import "fmt"

func main() {
    var k, v int
    fmt.Scan(&k)

    total := 0

    for i := 0; i < 5; i++ {
        fmt.Scan(&v)
        total += v

        fmt.Println(total >= k)
    }
}
