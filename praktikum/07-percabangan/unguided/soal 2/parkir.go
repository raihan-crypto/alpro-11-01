package main

import "fmt"

func main() {
	var jenis string
	var jam int
	fmt.Scan(&jenis, &jam)

	var total int
	if jenis == "Mobil" {
		if jam <= 1 {
			total = 5000
		} else {
			total = 5000 + (jam-1)*3000
		}
		if jam > 24 {
			total += 50000
		}
	} else if jenis == "Motor" {
		if jam <= 1 {
			total = 3000
		} else {
			total = 3000 + (jam-1)*2000
		}
		if jam > 24 {
			total += 30000
		}
	}

	fmt.Println(total)
}

