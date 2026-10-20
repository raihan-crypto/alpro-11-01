package main

import "fmt"

func main() {

	var nilai float64
	fmt.Scan(&nilai)

	if nilai >= 85 && nilai <= 100 {
		fmt.Println("A")
	} else if nilai <= 80 && nilai >70 {
		fmt.Println("B")
	} else if nilai <= 70 && nilai >60 {
		fmt.Println("C")
	} else if nilai <= 50 && nilai >0 {
		fmt.Println("E")
	}else{
		fmt.Println("Nilai tidak valid")
	}

}

// 85 - 100 A
// 70 - 80 B
// 60 - 70 C
// 50 - 60 D
// < 50 E
