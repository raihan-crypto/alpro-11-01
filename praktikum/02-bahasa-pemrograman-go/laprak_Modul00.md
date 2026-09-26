# <h1 align="center">Laporan Praktikum Modul 02 - Variabel, Tipe Data, dan Operasi</h1>
<p align="center">Raihan Ahmad Rahmadhani - [109092600025]</p>

## Dasar Teori

### A. Variabel dan Tipe Data di Go
Di bahasa Go, setiap variabel memiliki tipe data yang ketat (*strongly typed*). Beberapa tipe data dasar yang sering dipakai antara lain:
- **`int`**: Untuk bilangan bulat (misal: jumlah barang, skor).
- **`float64`**: Untuk bilangan desimal atau pecahan (misal: suhu, jari-jari lingkaran).
- **`string`**: Untuk teks/karakter (misal: nama).

Deklarasi variabel bisa ditulis eksplisit dengan `var namaVariabel tipeData` atau menggunakan *short declaration* `:=` jika nilainya langsung diisi pada saat pembuatan variabel.

### B. Input dan Output (I/O)
1. **Input (`fmt.Scan`)**: Mengambil masukan dari keyboard. Di Go, kita perlu menambahkan simbol `&` di depan variabel (seperti `fmt.Scan(&a)`) agar nilai masukan langsung disimpan ke alamat memori variabel tersebut.
2. **Output (`fmt.Println` & `fmt.Printf`)**: `fmt.Println` digunakan untuk mencetak teks biasa dengan baris baru otomatis, sedangkan `fmt.Printf` dipakai saat kita butuh format khusus menggunakan format *specifier* seperti `%d` (integer) atau `%g` / `%.2f` (desimal).

## Guided

### 1. skor.go

```go
package main

import "fmt"

func main() {
    var nama string
    var skormtk, skorbing int

    fmt.Scan(&nama)
    fmt.Scan(&skormtk)
    fmt.Scan(&skorbing)

    total := skormtk + skorbing
    
    ratarata := total / 2

    fmt.Println("Nama:", nama)
    fmt.Println("Total:", total)
    fmt.Println("Rata rata:", ratarata)
}
```
#### Deskripsi
Program ini membaca input berupa string (nama) dan dua nilai ujian (matematika & b.inggris). Program menjumlahkan kedua nilai tersebut menjadi variabel `total`, lalu menghitung `ratarata` dengan pembagian integer (`/ 2`). Hasilnya ditampilkan langsung ke layar.

### 2. tukar.go

```go
package main

import "fmt"

func main() {
    var a, b int

    fmt.Scan(&a, &b)

    a, b = b, a

    fmt.Printf("A = %d\n", a)
    fmt.Printf("B = %d\n", b)
}
```
#### Deskripsi
Program ini menukar nilai dari dua variabel `a` dan `b`. Di Go, pertukaran nilai bisa langsung dilakukan menggunakan fitur *multiple assignment* `a, b = b, a` tanpa perlu variabel bantuan (*temp*). Output dicetak menggunakan `fmt.Printf`.

### 3. lingkaran.go

```go
package main

import "fmt"

func main() {
    var r float64

    fmt.Scan(&r)

    phi := 3.14
    luas := phi * r * r

    fmt.Println("Hasil :", luas)
}
```
#### Deskripsi
Program ini menghitung luas lingkaran berdasarkan jari-jari `r` yang diinputkan pengguna. Karena melibatkan bilangan pecahan, tipe data yang digunakan adalah `float64` dengan konstanta `phi = 3.14`. Hasil perkalian `phi * r * r` kemudian ditampilkan ke terminal.

### 4. suhu.go

```go
package main

import "fmt"

func main() {
    var celsius float64

    fmt.Scan(&celsius)

    reamur := celsius * 4.0 / 5.0
    fahrenheit := celsius * 9.0 / 5.0 + 32.0
    kelvin := celsius + 273.15

    fmt.Printf("%g R, %g F, %g K\n", reamur, fahrenheit, kelvin)
}
```
#### Deskripsi
Program ini melakukan konversi suhu dari Celsius ke Reamur, Fahrenheit, dan Kelvin. Supaya perhitungannya presisi dan tidak terpotong aturan pembagian integer di Go, pecahan ditulis dalam bentuk desimal seperti `4.0 / 5.0` dan `9.0 / 5.0`. Output ditampilkan rapi dalam satu baris dengan format `%g`.

## Unguided

### 1. cacahuang.go

```go
package main

import "fmt"

func main() {
	var uang int

	fmt.Println("=== Program Pemecah Pecahan Uang ===")
	fmt.Print("Masukkan nominal uang: ")
	fmt.Scan(&uang)

	totalAwal := uang

	sepuluhRibu := uang / 10000
	uang = uang % 10000

	limaRibu := uang / 5000
	uang = uang % 5000

	seribu := uang / 1000

	fmt.Printf("\nHasil Pemecahan Uang Rp %d:\n", totalAwal)
	fmt.Printf("- Pecahan 10.000 : %d lembar\n", sepuluhRibu)
	fmt.Printf("- Pecahan 5.000  : %d lembar\n", limaRibu)
	fmt.Printf("- Pecahan 1.000  : %d lembar\n", seribu)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/raihan-crypto/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacah%20uang/output.png)

#### Deskripsi
Program ini memecah nominal uang ke dalam pecahan 10.000, 5.000, dan 1.000 rupiah. Logikanya menggunakan pembagian `/` untuk menghitung jumlah lembar pada pecahan terbesar dulu, dan modulus `%` untuk mengambil sisa uang yang belum terpecah. Program juga sudah dilengkapi pesan input dan rincian output yang jelas agar mudah dipahami pengguna.

### 2. kalkulator.go

```go
package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("=== Program Kalkulator Sederhana ===")
	fmt.Print("Masukkan bilangan bulat pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan bulat kedua   : ")
	fmt.Scan(&b)

	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	mod := a % b

	fmt.Println("\n=== Hasil Perhitungan ===")
	fmt.Printf("Penjumlahan (%d + %d) = %d\n", a, b, tambah)
	fmt.Printf("Pengurangan (%d - %d) = %d\n", a, b, kurang)
	fmt.Printf("Perkalian   (%d * %d) = %d\n", a, b, kali)
	fmt.Printf("Pembagian   (%d / %d) = %d\n", a, b, bagi)
	fmt.Printf("Sisa Bagi   (%d %% %d) = %d\n", a, b, mod)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/raihan-crypto/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/Kalkulator/output.png)

#### Deskripsi
Program kalkulator sederhana ini menerima dua angka input bertipe integer, lalu langsung mengeksekusi 5 operasi dasar matematika: tambah (`+`), kurang (`-`), kali (`*`), bagi (`/`), dan sisa bagi (`%`). Hasil perhitungannya dicetak baris per baris beserta format rumusnya menggunakan `fmt.Printf`.

## Kesimpulan
Pada praktikum Modul 02 ini, saya mempelajari bagaimana Go mengelola tipe data statis, deklarasi variabel (`var` dan `:=`), serta operasi input/output dasar. Poin penting yang saya catat adalah keharusan menggunakan pointer `&` saat `fmt.Scan`, kepraktisan fitur *multiple assignment* untuk swap variabel, serta pentingnya memperhatikan tipe data (integer vs float) dalam operasi aritmatika agar hasil perhitungan tidak terpotong.


