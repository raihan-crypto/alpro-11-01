package main

import "fmt"

func main() {
    var angka string
    fmt.Scan(&angka)

    a := int(angka[0] - '0')
    b := int(angka[1] - '0')
    c := int(angka[2] - '0')
    d := int(angka[3] - '0')

    terbesar := a
    terkecil := a

    if b > terbesar {
        terbesar = b
    }
    if c > terbesar {
        terbesar = c
    }
    if d > terbesar {
        terbesar = d
    }

    if b < terkecil {
        terkecil = b
    }
    if c < terkecil {
        terkecil = c
    }
    if d < terkecil {
        terkecil = d
    }

    fmt.Println(terbesar, terkecil)
}
