package main

import "fmt"

func main() {
    var nama string
    var skormtk, skorbing int

    fmt.Scan(&nama)
    fmt.Scan(&skormtk)
    fmt.Scan(&skorbing)

    total := skormtk + skorbing
    
    ratarata := total / 2

    fmt.Println("Nama:", nama)
 
    fmt.Println("Total:", total)
    fmt.Println("Rata rata:", ratarata)
}