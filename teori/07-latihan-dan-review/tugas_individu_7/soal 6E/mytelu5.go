package main

import "fmt"

func main() {
	countGold := 0
	countPlatinum := 0
	countSilver := 0

	totalPoinGold := 0
	totalPoinPlatinum := 0
	totalPoinSilver := 0

	totalPoin := 0

	// Berakhir apabila total poin keseluruhan pengguna telah mencapai 500
	for totalPoin < 500 {
		var poin int
		fmt.Scan(&poin)

		// Validasi: masukan harus >= 50
		for poin < 50 {
			fmt.Scan(&poin)
		}

		totalPoin += poin

		if poin > 200 {
			countGold++
			totalPoinGold += poin
		} else if poin >= 100 {
			countPlatinum++
			totalPoinPlatinum += poin
		} else {
			countSilver++
			totalPoinSilver += poin
		}
	}

	var rataGold, rataPlatinum, rataSilver float64

	if countGold > 0 {
		rataGold = float64(totalPoinGold) / float64(countGold)
	}
	if countPlatinum > 0 {
		rataPlatinum = float64(totalPoinPlatinum) / float64(countPlatinum)
	}
	if countSilver > 0 {
		rataSilver = float64(totalPoinSilver) / float64(countSilver)
	}

	fmt.Println(rataGold, rataPlatinum, rataSilver)
}

