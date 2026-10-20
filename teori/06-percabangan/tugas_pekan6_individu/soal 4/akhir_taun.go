package main

import "fmt"

func main() {
    var belanja int
    var bersedia bool

    fmt.Scan(&belanja, &bersedia)

    kartu := bersedia
    dapatDiskon := false
    dapatCashback := false

    diskon := 0
    cashback := 0

    if belanja >= 100000 {
        diskon = belanja * 10 / 100
        dapatDiskon = true
    }

    bayar := belanja - diskon

    if belanja >= 200000 && kartu {
        cashback = 75000
        dapatCashback = true
    }

    bayar = bayar - cashback

    fmt.Println("Kartu?", kartu)
    fmt.Println("Diskon?", dapatDiskon)
    fmt.Println("Cashback?", dapatCashback)
    fmt.Println("Rp.", bayar)
}
