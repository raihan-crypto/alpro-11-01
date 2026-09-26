package main

import "fmt"

func main() {
    var r float64

    fmt.Scan(&r)

    phi := 3.14
    luas := phi * r * r

    fmt.Println("Hasil :", luas)
}