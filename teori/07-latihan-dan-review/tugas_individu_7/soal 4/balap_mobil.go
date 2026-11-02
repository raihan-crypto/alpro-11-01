package main

import (
	"fmt"
	"strings"
)

func main() {
	var urutan []string
	for len(urutan) < 10 {
		var token string
		if _, err := fmt.Scan(&token); err != nil {
			break
		}
		for _, ch := range token {
			s := strings.ToUpper(string(ch))
			if s == "A" || s == "B" {
				urutan = append(urutan, s)
			}
		}
	}

	if len(urutan) != 10 {
		fmt.Println("tidak valid")
		return
	}

	// Hitung total mobil tim A dan B
	countA := 0
	countB := 0
	for _, m := range urutan {
		if m == "A" {
			countA++
		} else if m == "B" {
			countB++
		}
	}

	// Validasi: jumlah mobil tim A dan B masing-masing harus 5
	if countA != 5 || countB != 5 {
		fmt.Println("tidak valid")
		return
	}

	// Tentukan pemenang: tim pertama yang 4 mobilnya sampai garis finish
	finishA := 0
	finishB := 0
	pemenang := ""

	for _, m := range urutan {
		if m == "A" {
			finishA++
			if finishA == 4 {
				pemenang = "A"
				break
			}
		} else if m == "B" {
			finishB++
			if finishB == 4 {
				pemenang = "B"
				break
			}
		}
	}

	fmt.Println(pemenang)
}

