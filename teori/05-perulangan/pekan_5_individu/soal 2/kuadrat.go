package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    total := 0
    pertama := true

    for x > 0 {
        digit := x % 10
        kuadrat := digit * digit

        if !pertama {
            fmt.Print(" ")
        }

        fmt.Print(kuadrat)
        pertama = false

        total += kuadrat
        x = x / 10
    }

    fmt.Println()
    fmt.Println(total)
}
