package main

import "fmt"

func main() {
    var sebelumnya, sekarang float64

    fmt.Scan(&sebelumnya, &sekarang)

    if sekarang > sebelumnya {
        selisih := sekarang - sebelumnya
        fmt.Println("Peningkatan sebesar", selisih)
    } else if sekarang < sebelumnya {
        selisih := sebelumnya - sekarang
        fmt.Println("Penurunan sebesar", selisih)
    } else {
        fmt.Println("Tetap")
    }
}
