package main

import "fmt"

func main() {
    var h1, m1, h2, m2 int

    fmt.Scan(&h1, &m1, &h2, &m2)

    if h2 < h1 {
        h2 = h2 + 12
    }

    mulai := h1*60 + m1
    selesai := h2*60 + m2

    durasi := selesai - mulai

    hh := durasi / 60
    mm := durasi % 60

    fmt.Println(hh, "jam", mm, "menit")
}
