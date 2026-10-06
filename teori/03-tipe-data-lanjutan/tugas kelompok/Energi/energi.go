package main

import "fmt"

func main() {
	var E0, c, E1 int

	fmt.Scan(&E0, &c, &E1)

	// Energi yang terpakai = E0 - E1
	// cc = jumlah penggunaan penuh sebesar c dari energi yang terpakai
	cc := (E0 - E1) / c

	fmt.Println(cc)
}
