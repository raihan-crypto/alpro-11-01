package main

import "fmt"

func main() {
    var r float64

    fmt.Scan(&r)

    pi := 3.14
    luas := pi * r * r

    fmt.Println("Hasil :", luas)
}