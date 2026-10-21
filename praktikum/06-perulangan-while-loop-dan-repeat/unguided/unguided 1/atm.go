package main

import "fmt"

func main() {
	kesempatan := 3
	loginSukses := false
	var password string

	for kesempatan > 0 && !loginSukses {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == "1234" {
			loginSukses = true
		} else {
			kesempatan--
			fmt.Println("Password salah. Silakan coba lagi.")
		}
	}

	if loginSukses {
		fmt.Println("Login berhasil!")

		saldo := 10000000
		pilihan := 0

		for pilihan != 3 {
			fmt.Println("Menu ATM:")
			fmt.Println("1. Lihat Saldo")
			fmt.Println("2. Tarik Uang")
			fmt.Println("3. Keluar")
			fmt.Print("Pilih: ")
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				fmt.Printf("Saldo Anda: Rp %d\n", saldo)
			} else if pilihan == 2 {
				var tarik int
				fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
				fmt.Scan(&tarik)

				if tarik <= saldo {
					saldo -= tarik
					fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", tarik)
					fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
				} else {
					fmt.Println("Saldo tidak mencukupi.")
				}
			} else if pilihan == 3 {
				fmt.Println("Keluar dari sistem. Terima kasih!")
			}
		}
	} else {
		fmt.Println("ATM terblokir. Program selesai.")
	}
}
