package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Print("Masukkan tiga bilangan (n a b): ")
	fmt.Scan(&n, &a, &b)

	fmt.Printf("Kelipatan persekutuan %d dan %d : %v\n", a, b, n%a == 0 && n%b == 0)
}
