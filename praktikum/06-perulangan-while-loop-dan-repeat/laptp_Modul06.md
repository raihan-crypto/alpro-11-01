# <h1 align="center">Tugas Pendahuluan Modul 06 - Perulangan While</h1>
<p align="center">Raihan Ahmad Rahmadhani - 109092600025</p>

### 1. lulus.go

```go
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
```

##### Output

![Screenshot Output TP Soal 1](tp/soal%201/output.png)

#### Deskripsi
Program `lulus.go` ini dibuat untuk mengecek kelulusan nilai ujian siswa secara berulang-ulang sampai pengguna menginputkan angka `-1`.

Cara kerja programnya:
1. Program pertama kali meminta input nilai siswa dengan tampilan teks `Masukkan nilai Anda: `.
2. Karena di Go tidak ada kata kunci `while`, perulangan kondisi ditulis memakai `for nilai != -1`. Jadi selama nilai yang dimasukkan bukan `-1`, perulangan akan terus berjalan.
3. Di dalam perulangan, program mengecek nilai dengan `if`:
   - Jika nilainya 60 ke atas (`nilai >= 60`), program mencetak `"Anda lulus ujian."`.
   - Jika nilainya di bawah 60, program mencetak `"Anda tidak lulus ujian."`.
4. Setelah itu, program meminta input nilai baru lagi sebelum kembali mengecek kondisi di atas.
5. Saat pengguna memasukkan angka `-1`, perulangan berhenti dan program mencetak `"Program selesai."`.

---

### 2. deret.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := 0
	i := 1
	for i <= n {
		s += i
		i++
	}

	fmt.Println(s)
}
```

##### Output

![Screenshot Output TP Soal 2](tp/soal%202/output.png)

#### Deskripsi
Program `deret.go` ini dibuat untuk menghitung jumlah total deret angka dari 1 sampai batas `n` ($S = 1 + 2 + \dots + n$) menggunakan perulangan while.

Cara kerja programnya:
1. Program membaca satu input bilangan bulat positif `n` dari pengguna.
2. Kita menyiapkan variabel `s = 0` untuk menampung hasil penjumlahan deret dan variabel `i = 1` sebagai angka pencacah awalnya.
3. Perulangan while ditulis dengan `for i <= n`. Selama nilai `i` belum melebihi `n`, perulangan akan terus berputar.
4. Di setiap putaran, angka `i` ditambahkan ke dalam `s` (`s += i`), lalu nilai `i` dinaikkan satu dengan `i++`.
5. Begitu nilai `i` sudah lebih dari `n`, perulangan berhenti dan program mencetak hasil total `s`.

---

### 3. duaterendah.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var namaTerendah, namaKedua string
	var nilaiTerendah, nilaiKedua int

	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if i == 0 {
			namaTerendah = nama
			nilaiTerendah = nilai
		} else if i == 1 {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah
				namaTerendah = nama
				nilaiTerendah = nilai
			} else {
				namaKedua = nama
				nilaiKedua = nilai
			}
		} else {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah
				namaTerendah = nama
				nilaiTerendah = nilai
			} else if nilai < nilaiKedua {
				namaKedua = nama
				nilaiKedua = nilai
			}
		}
		i++
	}

	fmt.Println(namaTerendah)
	fmt.Println(namaKedua)
}
```

##### Output

![Screenshot Output TP Soal 3](tp/soal%203/output.png)

#### Deskripsi
Program `duaterendah.go` ini dibuat untuk mencari nama dua orang siswa yang memiliki nilai paling rendah dari total `n` siswa yang dimasukkan.

Cara kerja programnya:
1. Program membaca jumlah siswa `n`, lalu menyiapkan variabel untuk menyimpan nama dan nilai terendah pertama serta terendah kedua.
2. Perulangan while dijalankan sebanyak `n` kali dengan `i := 0` dan kondisi `for i < n`. Di setiap putaran dibaca pasangan `nama` dan `nilai` siswa.
3. Untuk menentukan urutannya:
   - Pada siswa pertama (`i == 0`), datanya langsung dicatat sebagai nilai terendah pertama.
   - Pada siswa kedua (`i == 1`), nilainya dibandingkan dengan siswa pertama untuk menentukan siapa yang di posisi pertama dan kedua.
   - Pada siswa berikutnya (`i >= 2`), jika ada nilai yang lebih kecil dari terendah pertama, data terendah pertama lama digeser ke posisi kedua, dan data baru menempati posisi pertama. Jika nilainya hanya lebih kecil dari posisi kedua, maka posisi kedua saja yang diperbarui.
4. Nilai `i` dinaikkan dengan `i++` di akhir loop agar perulangan terus maju.
5. Setelah semua data siswa selesai dicek, program mencetak nama siswa dengan nilai terendah dan kedua terendah.

---

### 4. biaya.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0
	tahun := 1

	for tahun < n {
		biaya = biaya * 1.06
		tahun++
	}

	fmt.Printf("%.0f\n", biaya)
}
```

##### Output

![Screenshot Output TP Soal 4](tp/soal%204/output.png)

#### Deskripsi
Program `biaya.go` ini berfungsi untuk menghitung perkiraan biaya kuliah pada tahun ke-`n`. Biaya awal di tahun pertama adalah Rp 10.000.000 dan setiap tahunnya bertambah 6% dari biaya tahun sebelumnya.

Cara kerja programnya:
1. Program membaca angka `n` yang menyatakan tahun yang ingin dicari biayanya.
2. Variabel `biaya` dibuat menggunakan tipe `float64` dengan nilai awal `10000000.0`, dan variabel `tahun` dimulai dari angka `1`.
3. Perulangan while ditulis dengan `for tahun < n`.
4. Di setiap putaran, biaya dikalikan `1.06` (karena naik 6%), lalu nilai `tahun` ditambah 1 (`tahun++`).
5. Ketika nilai `tahun` sudah sama dengan `n`, perulangan berhenti.
6. Hasil biaya dicetak menggunakan `fmt.Printf("%.0f\n", biaya)` agar angkanya dibulatkan ke rupiah terdekat tanpa menampilkan koma desimal.

---

## Kesimpulan

Dari tugas pendahuluan modul 06 mengenai perulangan while di Go, beberapa hal yang saya pelajari:
1. Di bahasa Go ternyata tidak ada kata kunci `while` seperti di bahasa lain. Untuk membuat perulangan while, kita cukup menuliskan `for` yang hanya berisi kondisi saja (`for kondisi { ... }`).
2. Perulangan while sangat berguna kalau kita belum tahu pasti berapa kali program harus berulang, misalnya pada program `lulus.go` yang baru berhenti ketika pengguna menginputkan angka `-1`.
3. Kita harus selalu ingat untuk mengubah nilai variabel kondisi di dalam badan loop (misalnya dengan `i++` atau meminta input baru) supaya perulangan tidak macet (*infinite loop*).
4. Kita juga bisa memakai perulangan while untuk menghitung total deret angka, mencari nilai terkecil, maupun menghitung kenaikan angka berkala seperti pada simulasi biaya kuliah.
