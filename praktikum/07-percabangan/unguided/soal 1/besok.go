package main

import "fmt"

func main() {
	var d, m, y int
	fmt.Scan(&d, &m, &y)

	if m < 1 || m > 12 || y <= 0 {
		fmt.Println("Tidak valid")
		return
	}

	var maxDays int
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		maxDays = 31
	case 4, 6, 9, 11:
		maxDays = 30
	case 2:
		isKabisat := (y%4 == 0 && y%100 != 0) || (y%400 == 0)
		if isKabisat {
			maxDays = 29
		} else {
			maxDays = 28
		}
	}

	if d < 1 || d > maxDays {
		fmt.Println("Tidak valid")
		return
	}

	nextD := d + 1
	nextM := m
	nextY := y

	if nextD > maxDays {
		nextD = 1
		nextM++
		if nextM > 12 {
			nextM = 1
			nextY++
		}
	}

	fmt.Println("Valid")
	fmt.Printf("%d/%d/%d\n", nextD, nextM, nextY)
}

