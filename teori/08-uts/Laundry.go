package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Persiapan
	jumlahPelanggan := 0
	totalBerat := 0.0
	totalPendapatan := 0.0

	var namaList []string
	var idList []string
	var beratList []float64
	var layananList []int
	var biayaList []float64
	var diskonList []float64
	var totalList []float64

	reader := bufio.NewReader(os.Stdin)

	// Fungsi untuk membaca input
	bacaInput := func(pesan string) string {
		fmt.Print(pesan)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	fmt.Println("===================================")
	fmt.Println("       SISTEM LAUNDRY GO")
	fmt.Println("===================================")

	// Proses pelanggan
	for {
		// 2. Input nama
		nama := bacaInput("\nMasukkan nama pelanggan: ")

		for nama == "" {
			fmt.Println("Nama tidak boleh kosong.")
			nama = bacaInput("Masukkan nama pelanggan: ")
		}

		// 3. Membuat ID pelanggan (nama kembar diberi nomor)
		kembar := 0

		for i := 0; i < jumlahPelanggan; i++ {
			if strings.EqualFold(namaList[i], nama) {
				kembar++
			}
		}

		id := nama

		if kembar > 0 {
			id = fmt.Sprintf("%s %d", nama, kembar+1)
		}

		// 4. Input berat dan validasi
		var berat float64

		for {
			input := bacaInput("Masukkan berat cucian (kg): ")
			var err error

			berat, err = strconv.ParseFloat(input, 64)

			if err == nil && berat > 0 {
				break
			}

			fmt.Println("Berat tidak valid! Berat harus lebih dari 0.")
		}

		// 5. Input layanan dan validasi
		var layanan int

		for {
			fmt.Println("\nPilihan layanan:")
			fmt.Println("1. Reguler (Rp5.000/kg)")
			fmt.Println("2. Ekspres (Rp8.000/kg)")

			input := bacaInput("Pilih layanan (1/2): ")
			var err error

			layanan, err = strconv.Atoi(input)

			if err == nil && (layanan == 1 || layanan == 2) {
				break
			}

			fmt.Println("Pilihan salah! Silakan pilih 1 atau 2.")
		}

		// 6. Menentukan tarif per kg
		var tarif float64
		var namaLayanan string

		if layanan == 2 {
			tarif = 8000
			namaLayanan = "Ekspres"
		} else {
			tarif = 5000
			namaLayanan = "Reguler"
		}

		// 7. Menghitung biaya, diskon, dan total
		biaya := berat * tarif
		var diskon float64

		if berat >= 10 {
			diskon = 0.10 * biaya
		} else {
			diskon = 0
		}

		total := biaya - diskon

		// 8. Mencetak nota
		fmt.Println("\n===================================")
		fmt.Println("            NOTA LAUNDRY")
		fmt.Println("===================================")
		fmt.Println("ID Pelanggan :", id)
		fmt.Printf("Berat Cucian : %.2f kg\n", berat)
		fmt.Println("Layanan      :", namaLayanan)
		fmt.Printf("Biaya        : Rp%.2f\n", biaya)
		fmt.Printf("Diskon       : Rp%.2f\n", diskon)
		fmt.Printf("Total Bayar  : Rp%.2f\n", total)
		fmt.Println("===================================")

		// 9. Menyimpan nota dan memperbarui rekap
		jumlahPelanggan++

		namaList = append(namaList, nama)
		idList = append(idList, id)
		beratList = append(beratList, berat)
		layananList = append(layananList, layanan)
		biayaList = append(biayaList, biaya)
		diskonList = append(diskonList, diskon)
		totalList = append(totalList, total)

		totalBerat += berat
		totalPendapatan += total

		// 10. Menanyakan apakah ingin melanjutkan
		for {
			lagi := strings.ToLower(
				bacaInput("\nProses pelanggan lain? (y/t): "),
			)

			if lagi == "t" {
				goto SELESAI
			}

			if lagi == "y" {
				break
			}

			fmt.Println("Masukkan y untuk lanjut atau t untuk selesai.")
		}
	}

SELESAI:

	// 11. Menampilkan semua nota
	fmt.Println("\n===================================")
	fmt.Println("         DAFTAR SEMUA NOTA")
	fmt.Println("===================================")

	if jumlahPelanggan == 0 {
		fmt.Println("Belum ada pelanggan.")
	}

	for i := 0; i < jumlahPelanggan; i++ {
		fmt.Printf("\nNota ke-%d\n", i+1)
		fmt.Println("ID Pelanggan :", idList[i])
		fmt.Printf("Berat        : %.2f kg\n", beratList[i])

		if layananList[i] == 1 {
			fmt.Println("Layanan      : Reguler")
		} else {
			fmt.Println("Layanan      : Ekspres")
		}

		fmt.Printf("Biaya        : Rp%.2f\n", biayaList[i])
		fmt.Printf("Diskon       : Rp%.2f\n", diskonList[i])
		fmt.Printf("Total Bayar  : Rp%.2f\n", totalList[i])
	}

	// 12. Rekap akhir
	fmt.Println("\n===================================")
	fmt.Println("             REKAP AKHIR")
	fmt.Println("===================================")
	fmt.Println("Jumlah Pelanggan :", jumlahPelanggan)
	fmt.Printf("Total Berat      : %.2f kg\n", totalBerat)
	fmt.Printf("Total Pendapatan : Rp%.2f\n", totalPendapatan)

	// 13. Mencari nota berdasarkan ID
	for {
		cariID := bacaInput(
			"\nCari nota (ketik ID, atau x untuk keluar): ",
		)

		if strings.EqualFold(cariID, "x") {
			break
		}

		ketemu := false

		for i := 0; i < jumlahPelanggan; i++ {
			if strings.EqualFold(idList[i], cariID) {
				fmt.Println("\n===================================")
				fmt.Println("         NOTA DITEMUKAN")
				fmt.Println("===================================")
				fmt.Println("ID Pelanggan :", idList[i])
				fmt.Printf("Berat        : %.2f kg\n", beratList[i])

				if layananList[i] == 1 {
					fmt.Println("Layanan      : Reguler")
				} else {
					fmt.Println("Layanan      : Ekspres")
				}

				fmt.Printf("Biaya        : Rp%.2f\n", biayaList[i])
				fmt.Printf("Diskon       : Rp%.2f\n", diskonList[i])
				fmt.Printf("Total Bayar  : Rp%.2f\n", totalList[i])
				fmt.Println("===================================")

				ketemu = true
				break
			}
		}

		if !ketemu {
			fmt.Println("Nota tidak ditemukan.")
		}
	}

	fmt.Println("\nProgram laundry selesai. Terima kasih!")
}
