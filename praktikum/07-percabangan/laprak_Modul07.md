# <h1 align="center">Laporan Praktikum Modul 07 - Percabangan</h1>
<p align="center">Raihan Ahmad Rahmadhani - 109092600025</p>

## Dasar Teori

### A. Konsep Percabangan (*Selection / Branching*)
Percabangan (*branching/selection*) merupakan struktur kontrol mendasar dalam algoritma dan pemrograman yang memungkinkan komputer membuat keputusan dan memilih jalur eksekusi tertentu berdasarkan kondisi logika bernilai *boolean* (`true` atau `false`). Tanpa percabangan, seluruh instruksi dalam program hanya dapat berjalan secara sekuensial (lurus dari baris awal hingga akhir).

### B. Operator Relasional dan Logika dalam Go
Evaluasi kondisi percabangan dibangun atas kombinasi operator relasional dan operator logika:
1. **Operator Relasional**: Membandingkan dua operan, meliputi `==` (sama dengan), `!=` (tidak sama dengan), `<` (kurang dari), `>` (lebih dari), `<=` (kurang dari sama dengan), dan `>=` (lebih dari sama dengan).
2. **Operator Logika**:
   - `&&` (AND): Bernilai `true` hanya apabila kedua kondisi bernilai `true`.
   - `||` (OR): Bernilai `true` apabila salah satu atau kedua kondisi bernilai `true`.
   - `!` (NOT): Membalikkan nilai logika.
3. **Short-Circuit Evaluation**: Go mengevaluasi ekspresi logika secara efisien. Pada ekspresi `A && B`, jika `A` bernilai `false`, maka `B` tidak akan dievaluasi. Pada `A || B`, jika `A` bernilai `true`, maka `B` tidak akan dievaluasi.

### C. Struktur `if`, `if-else`, dan `else if`
1. **`if` (Satu Kondisi)**: Blok kode hanya dieksekusi jika kondisi bernilai `true`.
2. **`if-else` (Dua Alternatif)**: Memilih salah satu dari dua blok instruksi; jika kondisi `if` bernilai `false`, maka blok `else` akan dijalankan.
3. **`else if` (Banyak Alternatif)**: Menguji sekumpulan kondisi secara beruntun (*cascade*). Kondisi pertama yang bernilai `true` akan dieksekusi, dan sisa cabang di bawahnya diabaikan.

### D. Struktur `switch - case` pada Go
Pernyataan `switch` menyederhanakan penulisan banyak kondisi yang menguji nilai spesifik (*discrete values*):
1. **Implicit Break**: Berbeda dengan bahasa C/C++ atau Java, setiap `case` pada Go secara otomatis memiliki sifat *break*. Aliran eksekusi tidak akan merembes (*fall-through*) ke `case` berikutnya kecuali jika kata kunci `fallthrough` dituliskan secara eksplisit.
2. **Tagless Switch**: Go mendukung penulisan `switch` tanpa variabel di kepalanya (`switch { ... }`), di mana setiap `case` dapat berupa ekspresi boolean independen layaknya rantai `else if`.
3. **Multiple Values per Case**: Satu `case` di Go dapat memuat beberapa nilai sekaligus yang dipisahkan oleh tanda koma (contoh: `case 1, 3, 5, 7, 8, 10, 12:`).

---

## Guided

### 1. kabisat.go

```go
package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
```

#### Deskripsi
Program `kabisat.go` bertujuan untuk menentukan apakah suatu tahun yang diinputkan pengguna merupakan tahun kabisat atau bukan. 
Ketentuan tahun kabisat:
1. Habis dibagi 4 dan tidak habis dibagi 100 (`tahun % 4 == 0 && tahun % 100 != 0`), ATAU
2. Habis dibagi 400 (`tahun % 400 == 0`).

Alur kerja program:
- Program membaca sebuah bilangan bulat `tahun` dari pengguna.
- Mengevaluasi kondisi kabisat menggunakan struktur `if-else`.
- Jika kondisi terpenuhi, program mencetak `"Ya"`. Jika tidak, program mencetak `"Tidak"`.

---

### 2. segitiga.go

```go
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan segitiga")
	} else if a == b && b == c {
		fmt.Println("Segitiga sama sisi")
	} else if a == b || a == c || b == c {
		fmt.Println("Segitiga sama kaki")
	} else {
		fmt.Println("Segitiga sembarang")
	}
}
```

#### Deskripsi
Program `segitiga.go` menerima panjang tiga buah sisi segitiga (`a`, `b`, dan `c`) lalu mengklasifikasikan jenis segitiga tersebut.
Alur penentuan:
1. **Validasi Geometri Segitiga**: Sebuah segitiga valid hanya jika seluruh panjang sisinya bernilai positif (> 0) serta memenuhi pertidaksamaan segitiga (jumlah dua sisi mana pun harus lebih besar dari sisi ketiga). Jika `a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a`, program langsung mencetak `"Bukan segitiga"`.
2. **Segitiga Sama Sisi**: Jika ketiga sisi sama panjang (`a == b && b == c`), program mencetak `"Segitiga sama sisi"`.
3. **Segitiga Sama Kaki**: Jika tepat dua sisi bernilai sama (`a == b || a == c || b == c`), dicetak `"Segitiga sama kaki"`.
4. **Segitiga Sembarang**: Jika ketiga sisinya memiliki panjang yang berbeda-beda, dicetak `"Segitiga sembarang"`.

---

### 3. kalkulator.go

```go
package main

import "fmt"

func main() {
	var a, b int
	var op string
	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
		} else {
			fmt.Printf("%.2f\n", float64(a)/float64(b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}
```

#### Deskripsi
Program `kalkulator.go` mengimplementasikan operasi kalkulator sederhana berbasis struktur `switch-case` untuk memproses dua bilangan bulat (`a` dan `b`) dan satu operator aritmetika (`op`).
Ketentuan operasional:
- Operator `+`, `-`, `*`: Menghitung operasi aritmetika dasar dan mencetak hasil berupa bilangan bulat.
- Operator `/`: Mengonversi operan ke `float64` dan mencetak hasil pembagian dengan presisi 2 angka di belakang koma (`%.2f`). Jika pembagi bernilai 0 (`b == 0`), ditampilkan pesan kesalahan `"Error: pembagian dengan nol"`.
- Operator `%`: Menghitung sisa hasil bagi (*modulo*). Menampilkan pesan kesalahan jika `b == 0`.
- Blok `default`: Menangani operator di luar yang ditentukan dengan mencetak `"Operator tidak valid"`.

---

### 4. ongkir.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum (30 kg)")
		return
	}

	var tarifPerKg int
	switch zona {
	case 1:
		tarifPerKg = 10000
	case 2:
		tarifPerKg = 15000
	case 3:
		tarifPerKg = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}

	beratDihitung := int(math.Ceil(berat))
	total := beratDihitung * tarifPerKg

	if layanan == "Express" {
		total += total * 50 / 100
	}

	fmt.Println(beratDihitung)
	fmt.Println(total)
}
```

#### Deskripsi
Program `ongkir.go` digunakan untuk menghitung tarif pengiriman paket berdasarkan tiga parameter masukan: berat paket (desimal dalam kg), zona tujuan pengiriman (1, 2, atau 3), dan jenis layanan (`Reguler` atau `Express`).
Alur program:
1. Memvalidasi berat: jika `berat > 30`, program langsung mencetak `"Berat melebihi batas maksimum (30 kg)"` dan berhenti.
2. Menggunakan `switch zona` untuk menentukan tarif dasar per kg: Zona 1 (Rp 10.000), Zona 2 (Rp 15.000), Zona 3 (Rp 25.000). Jika zona di luar 1, 2, atau 3, program menampilkan `"Zona tidak valid"` dan keluar.
3. Membulatkan berat ke atas ke kilogram terdekat menggunakan fungsi `math.Ceil(berat)` dari pustaka `math`.
4. Menghitung biaya dasar: `total = beratDihitung * tarifPerKg`.
5. Apabila jenis layanan adalah `Express`, dikenakan biaya tambahan sebesar 50% (`total += total * 50 / 100`).
6. Mencetak nilai berat yang dihitung dan total ongkos kirim ke terminal.

---

## Unguided

### 1. besok.go

```go
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
```

##### Output
![Screenshot Output Unguided Soal 1](unguided/soal%201/output.png)

#### Deskripsi
Program `besok.go` menerima input tanggal (`d`), bulan (`m`), dan tahun (`y`). Program memvalidasi apakah tanggal tersebut ada dalam kalender Masehi, lalu menampilkan tanggal keesokan harinya jika valid.
Alur kerja:
1. Memeriksa validitas nomor bulan (`1 <= m <= 12`) dan tahun (`y > 0`).
2. Menentukan jumlah hari maksimum pada bulan tersebut menggunakan struktur `switch m`:
   - Bulan dengan 31 hari: 1, 3, 5, 7, 8, 10, 12.
   - Bulan dengan 30 hari: 4, 6, 9, 11.
   - Bulan 2 (Februari): 29 hari jika tahun kabisat, 28 hari jika tahun biasa.
3. Memeriksa validitas tanggal: jika `d < 1 || d > maxDays`, program mencetak `"Tidak valid"`.
4. Jika valid, program menghitung tanggal besok:
   - Jika `d < maxDays`, maka hari bertambah satu (`d + 1`).
   - Jika `d == maxDays`, maka hari direset ke 1 dan bulan bertambah satu (`m + 1`).
   - Jika bulan melebihi 12 (akhir tahun), bulan direset ke 1 dan tahun bertambah satu (`y + 1`).
5. Mencetak `"Valid"` dan tanggal besok dengan format `d/m/y`.

---

### 2. bmi.go 

> **Perhitungan NIM:**  
> NIM: `109092600025`  
> Nilai $r = 109092600025 \pmod 3 = 1$ $\rightarrow$ Mengerjakan **Soal B** (`bmi.go`).

```go
package main

import "fmt"

func main() {
	var berat, tinggiCm float64
	fmt.Scan(&berat, &tinggiCm)

	tinggiM := tinggiCm / 100.0
	bmi := berat / (tinggiM * tinggiM)

	var kategori string
	if bmi < 18.5 {
		kategori = "Kurus"
	} else if bmi < 25.0 {
		kategori = "Normal"
	} else if bmi < 30.0 {
		kategori = "Gemuk"
	} else {
		kategori = "Obesitas"
	}

	fmt.Printf("%.2f\n", bmi)
	fmt.Println(kategori)
}
```

##### Output
![Screenshot Output Unguided Soal 2](unguided/soal%202/output.png)

#### Deskripsi
Program `bmi.go` menghitung nilai *Body Mass Index* (BMI) berdasarkan masukan berat badan (kg) dan tinggi badan (cm), kemudian mengelompokkannya ke dalam kategori indeks massa tubuh.
Alur kerja:
1. Membaca masukan berat badan (`berat`) dan tinggi badan dalam centimeter (`tinggiCm`).
2. Mengonversi tinggi badan dari centimeter ke meter (`tinggiM = tinggiCm / 100.0`).
3. Menghitung BMI dengan rumus:
   $$\text{BMI} = \frac{\text{berat}}{\text{tinggiM} \times \text{tinggiM}}$$
4. Mengklasifikasikan kategori BMI menggunakan struktur `if - else if - else`:
   - $\text{BMI} < 18.5 \rightarrow \text{Kurus}$
   - $18.5 \le \text{BMI} < 25 \rightarrow \text{Normal}$
   - $25 \le \text{BMI} < 30 \rightarrow \text{Gemuk}$
   - $\text{BMI} \ge 30 \rightarrow \text{Obesitas}$
5. Menampilkan nilai BMI yang diformat dengan dua angka di belakang koma (`%.2f`) serta nama kategori berat badannya.

---

### Lampiran Tambahan: Soal Alternatif Unguided 2

Untuk kelengkapan referensi belajar praktikum, berikut adalah implementasi untuk Soal A dan Soal C:

#### Soal A: parkir.go
```go
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
```

#### Soal C: kereta.go
```go
package main

import "fmt"

func main() {
	var kelas string
	var usia int
	var hari string

	fmt.Scan(&kelas, &usia, &hari)

	var hargaDasar float64
	switch kelas {
	case "Ekonomi":
		hargaDasar = 100000
	case "Bisnis":
		hargaDasar = 200000
	case "Eksekutif":
		hargaDasar = 350000
	default:
		fmt.Println("Kelas tidak valid")
		return
	}

	harga := hargaDasar
	if hari == "Sabtu" || hari == "Minggu" {
		harga += harga * 0.20
	}

	if usia < 3 {
		harga = 0
	} else if usia >= 3 && usia <= 12 {
		harga -= harga * 0.50
	} else if usia >= 60 {
		harga -= harga * 0.30
	}

	fmt.Println(int(harga))
}
```

---

## Kesimpulan

Dari pelaksanaan Praktikum Modul 07 mengenai percabangan dalam bahasa pemrograman Go, dapat ditarik beberapa kesimpulan penting:
1. **Penerapan Struktur Percabangan yang Tepat**:
   - Struktur `if-else` sangat ideal digunakan untuk evaluasi kondisi yang kompleks, rentang nilai (*range*), pertidaksamaan geometri, serta kondisi gabungan operator logika (`&&`, `||`, `!`).
   - Struktur `switch-case` sangat efektif dan efisien untuk perbandingan nilai diskret konstan (seperti operator kalkulator, kode zona pengiriman, nomor bulan kalender, atau kelas tiket kereta) karena menghasilkan kode yang jauh lebih ringkas, terstruktur, dan mudah dibaca.
2. **Kelebihan Fitur Switch pada Go**:
   - *Implicit Break* pada Go mencegah terjadinya kesalahan logika akibat *fall-through* yang tidak disengaja.
   - Fleksibilitas pengelompokan banyak nilai dalam satu baris `case` (seperti pada daftar bulan) membuat penulisan logika lebih bersih dan hemat baris kode.
3. **Pentingnya Validasi Masukan dan Penanganan Kasus Khusus (*Edge Cases*)**:
   - Validasi masukan (seperti pemeriksaan pembagian dengan nol, batas berat maksimum paket, panjang sisi segitiga yang tidak memenuhi syarat, maupun jumlah hari spesifik untuk bulan kabisat) sangat penting dalam memastikan program berjalan kokoh (*robust*) dan tidak mengalami *runtime panic/crash*.
