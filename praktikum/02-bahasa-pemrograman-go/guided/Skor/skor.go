package main

import "fmt"

func main() {
    var nama string
    var skormtk, skorbing int

    fmt.Print("Nama: ")
    fmt.Scan(&nama)

    fmt.Print("Skor Mtk: ")
    fmt.Scan(&skormtk)

    fmt.Print("Skor B.ing: ")
    fmt.Scan(&skorbing)

    // Initialize variables with := and add values directly without &
    total := skormtk + skorbing
    
    // Convert to float64 to prevent integer truncation during division
    ratarata := float64(total) / 2.0

    fmt.Println("Nama :", nama)
    fmt.Println("Jumlah skor :", total)
    fmt.Println("Ratarata :", ratarata)
}