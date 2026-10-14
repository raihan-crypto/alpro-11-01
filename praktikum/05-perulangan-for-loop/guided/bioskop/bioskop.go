
package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)

    totalPendapatan := 0

    for i := 1; i <= t; i++ {
        var film, hari string
        var jumlahTiket int

        fmt.Scan(&film, &hari, &jumlahTiket)

        harga := 0

        if hari == "Senin" || hari == "Selasa" ||
            hari == "Rabu" || hari == "Kamis" {
            if film == "Reguler" {
                harga = 40000
            } else if film == "3D" {
                harga = 60000
            } else if film == "IMAX" {
                harga = 80000
            }
        } else {
            if film == "Reguler" {
                harga = 50000
            } else if film == "3D" {
                harga = 70000
            } else if film == "IMAX" {
                harga = 90000
            }
        }

        total := harga * jumlahTiket

        if jumlahTiket >= 5 {
            total = total - (total * 10 / 100)
        }

        fmt.Println(total)
        totalPendapatan = totalPendapatan + total
    }

    fmt.Println(totalPendapatan)
}
