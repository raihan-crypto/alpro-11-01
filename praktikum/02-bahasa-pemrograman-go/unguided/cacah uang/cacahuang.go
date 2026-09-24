package main

import "fmt"

func main() {
    var uang int

    fmt.Scan(&uang)

    sepuluhRibu := uang / 10000
    uang = uang % 10000

    limaRibu := uang / 5000
    uang = uang % 5000

    seribu := uang / 1000

    fmt.Println(sepuluhRibu, limaRibu, seribu)
}