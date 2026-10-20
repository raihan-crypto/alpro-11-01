package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai Anda: ")
	fmt.Scan(&nilai)

	for nilai != -1 {
		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}

		fmt.Print("Masukkan nilai Anda: ")
		fmt.Scan(&nilai)
	}

	fmt.Println("Program selesai.")
}
