package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    cantik := true
    digitKanan := x % 10
    x = x / 10

    for x > 0 {
        digitKiri := x % 10

        if digitKiri > digitKanan {
            cantik = false
            break
        }

        digitKanan = digitKiri
        x = x / 10
    }

    fmt.Println(cantik)
}
