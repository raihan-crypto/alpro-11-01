package main

import "fmt"

func main() {
    var p int
    fmt.Scan(&p)

    dewasa := 0
    kecil := 0
    tidakBerangkat := 0

    if p <= 15 {
        dewasa = (p + 4) / 5
    } else {
        dewasa = 3
        sisa := p - 15

        if sisa <= 10 {
            kecil = (sisa + 1) / 2
        } else {
            kecil = 5
            tidakBerangkat = sisa - 10
        }
    }

    fmt.Println("Dewasa", dewasa)

    if kecil > 0 {
        fmt.Println("Kecil", kecil)
    }

    if tidakBerangkat > 0 {
        fmt.Println("dan", tidakBerangkat, "tak berangkat")
    }
}
