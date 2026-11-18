# <h1 align="center">Tugas Pendahuluan Modul 10 - Fungsi</h1>
<p align="center">Raihan Ahmad Rahmadhani - 109092600025</p>

---

### 1. Konsep Fungsi dan Penggunaannya dalam Bahasa Go

#### Deskripsi

Fungsi (*function*) adalah blok kode yang dirancang untuk melakukan tugas tertentu dan dapat dipanggil berulang kali dari bagian lain program. Fungsi merupakan salah satu konsep dasar dalam pemrograman terstruktur yang membantu memecah program besar menjadi bagian-bagian kecil yang lebih mudah dipahami, dikelola, dan diuji.

Dalam bahasa Go, fungsi didefinisikan menggunakan kata kunci `func` dengan sintaks sebagai berikut:

```go
func namaFungsi(parameter tipe) tipePengembalian {
    // isi fungsi
    return nilai
}
```

**Komponen utama fungsi dalam Go:**

1. **Nama fungsi** — identifier yang digunakan untuk memanggil fungsi. Nama dimulai dengan huruf kapital jika ingin dapat diakses dari luar *package* (exported), atau huruf kecil jika hanya untuk penggunaan internal.

2. **Parameter** — nilai masukan yang diterima fungsi. Go mendukung nol atau lebih parameter, dan setiap parameter harus menyertakan tipenya secara eksplisit. Contoh: `func tambah(a int, b int) int`.

3. **Tipe kembalian (return type)** — tipe data dari nilai yang dikembalikan fungsi. Jika fungsi tidak mengembalikan nilai, tipe kembalian dapat dihilangkan. Go juga mendukung **multiple return values**, misalnya `func bagi(a, b int) (int, error)`.

4. **Pernyataan `return`** — digunakan untuk mengembalikan nilai dari fungsi. Jika fungsi bertipe `void` (tanpa kembalian), `return` bersifat opsional.

**Keunggulan penggunaan fungsi:**

- **Reusabilitas** — kode yang sama tidak perlu ditulis ulang, cukup panggil fungsinya.
- **Modularitas** — program terbagi menjadi modul-modul kecil yang mudah dikembangkan secara terpisah.
- **Keterbacaan** — program menjadi lebih mudah dibaca karena setiap fungsi memiliki tanggung jawab yang jelas.
- **Kemudahan debugging** — kesalahan lebih mudah dilacak karena terisolasi dalam satu fungsi.

**Fungsi dengan struct dan type alias:**

Dalam Go, fungsi sering digunakan bersama `struct` dan `type alias` untuk memproses data terstruktur. `type alias` memungkinkan kita memberi nama baru pada tipe yang sudah ada (misalnya array dari struct), sehingga kode menjadi lebih ekspresif dan mudah dibaca.

```go
type Mahasiswa struct {
    nama  string
    nilai int
}
type DaftarMahasiswa = [100]Mahasiswa

func prosesData(T DaftarMahasiswa, n int) float64 {
    // ...
}
```

Dengan konsep ini, fungsi tidak hanya memproses tipe data primitif, tetapi juga dapat menerima dan mengembalikan tipe data kompleks, menjadikan Go sangat fleksibel untuk membangun sistem yang terstruktur.

---

### 2. raport.go

```go
package main

import "fmt"

const NMAX = 999

type rapot struct {
	matkul string
	nilai  int
}

type tabRapot = [NMAX]rapot

func nilaiHuruf(x int) string {
	if x > 80 {
		return "A"
	} else if x > 70 {
		return "AB"
	} else if x > 65 {
		return "B"
	} else if x > 60 {
		return "BC"
	} else if x > 50 {
		return "C"
	} else if x > 40 {
		return "D"
	} else {
		return "E"
	}
}

func rataRata(T tabRapot, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].nilai
	}
	return float64(total) / float64(n)
}

func main() {
	var n int
	var T tabRapot
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].matkul, &T[i].nilai)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %s\n", T[i].matkul, T[i].nilai, nilaiHuruf(T[i].nilai))
	}
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}
```

##### Output

![Screenshot Output TP Soal 1](TP/soal%201/Screenshot%202026-11-18%20at%2013.56.55.png)

#### Deskripsi

Program `raport.go` merupakan program cetak rapor mahasiswa yang memanfaatkan fungsi untuk memisahkan logika konversi nilai dan perhitungan rata-rata dari fungsi utama `main`.

1. Program mendeklarasikan konstanta `NMAX = 999` sebagai batas maksimum data yang dapat ditampung.
2. Tipe `rapot` didefinisikan sebagai `struct` yang menyimpan dua field: `matkul` (nama mata kuliah, bertipe `string`) dan `nilai` (nilai numerik, bertipe `int`).
3. `tabRapot` merupakan *type alias* dari array berukuran `NMAX` bertipe `rapot`, digunakan sebagai tipe parameter pada fungsi-fungsi yang memproses data rapor.
4. Fungsi `nilaiHuruf(x int) string` menerima nilai numerik dan mengembalikan nilai huruf berdasarkan rentang nilai: di atas 80 → `"A"`, di atas 70 → `"AB"`, di atas 65 → `"B"`, di atas 60 → `"BC"`, di atas 50 → `"C"`, di atas 40 → `"D"`, selain itu → `"E"`.
5. Fungsi `rataRata(T tabRapot, n int) float64` menjumlahkan seluruh nilai dari `n` data pertama dalam array `T`, lalu membaginya dengan `n` untuk menghasilkan nilai rata-rata bertipe `float64`.
6. Pada `main`, program membaca jumlah mata kuliah `n`, lalu membaca pasangan `matkul` dan `nilai` sebanyak `n` kali ke dalam array `T`.
7. Program kemudian mencetak setiap baris rapor dalam format `matkul nilai nilaiHuruf` menggunakan `fmt.Printf`, dengan nilai huruf diperoleh dari pemanggilan fungsi `nilaiHuruf`.
8. Terakhir, program mencetak rata-rata nilai menggunakan pemanggilan fungsi `rataRata`, diformat dua angka di belakang koma.

---

### 3. inputnilai.go

```go
package main

import "fmt"

const NMAX = 100

type student struct {
	nim   string
	name  string
	grade float64
}

type students = [NMAX]student

func cariNIM(T students, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}

func rataRata(T students, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}

func main() {
	var T students
	var n int
	var nim, name string
	var grade float64

	for {
		fmt.Scan(&nim)
		if nim == "STOP" {
			break
		}
		fmt.Scan(&name, &grade)
		if cariNIM(T, n, nim) == -1 {
			T[n] = student{nim, name, grade}
			n++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(T[i].nim, T[i].name, T[i].grade)
	}
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}
```

##### Output

![Screenshot Output TP Soal 2](TP/soal%202/Screenshot%202026-11-18%20at%2013.58.17.png)

#### Deskripsi

Program `inputnilai.go` merupakan sistem input data nilai mahasiswa yang dilengkapi dengan validasi duplikasi NIM menggunakan fungsi pencarian, serta perhitungan rata-rata nilai akhir.

1. Program mendeklarasikan konstanta `NMAX = 100` sebagai batas maksimum jumlah mahasiswa yang dapat disimpan.
2. Tipe `student` didefinisikan sebagai `struct` dengan tiga field: `nim` (Nomor Induk Mahasiswa, bertipe `string`), `name` (nama mahasiswa, bertipe `string`), dan `grade` (nilai akhir, bertipe `float64`).
3. `students` merupakan *type alias* dari array berukuran `NMAX` bertipe `student`, digunakan sebagai tipe parameter pada kedua fungsi pembantu.
4. Fungsi `cariNIM(T students, n int, nim string) int` melakukan pencarian linier pada `n` data pertama array `T`. Jika NIM ditemukan, fungsi mengembalikan indeks posisinya; jika tidak ditemukan, fungsi mengembalikan `-1`.
5. Fungsi `rataRata(T students, n int) float64` menjumlahkan nilai `grade` dari seluruh `n` mahasiswa, lalu membaginya dengan `n` untuk menghasilkan rata-rata bertipe `float64`.
6. Pada `main`, program membaca NIM satu per satu dalam *loop* tak terbatas (`for` tanpa kondisi). Jika NIM yang dibaca adalah `"STOP"`, *loop* dihentikan dengan `break`.
7. Untuk setiap NIM yang bukan `"STOP"`, program membaca `name` dan `grade`, kemudian memanggil `cariNIM` untuk memeriksa duplikasi. Data hanya disimpan ke array `T` jika NIM belum ada (nilai kembalian `cariNIM` adalah `-1`), sehingga mencegah entri ganda.
8. Setelah input selesai, program mencetak seluruh data mahasiswa yang tersimpan menggunakan `fmt.Println`.
9. Program mengakhiri outputnya dengan mencetak rata-rata nilai menggunakan pemanggilan fungsi `rataRata`, diformat dua angka di belakang koma.

---

## Kesimpulan

1. **Fungsi sebagai unit modular** — Penggunaan fungsi seperti `nilaiHuruf` dan `cariNIM` memisahkan logika spesifik dari alur utama program, membuat kode lebih terstruktur, mudah dibaca, dan mudah dimodifikasi tanpa memengaruhi bagian lain.

2. **Type alias dan struct memperkuat fungsi** — Kombinasi `struct` untuk mengelompokkan data terkait dan `type alias` untuk mendefinisikan tipe array kustom memungkinkan fungsi menerima dan memproses data terstruktur secara ekspresif dan aman secara tipe (*type-safe*).

3. **Fungsi mencegah duplikasi kode** — Fungsi `rataRata` pada kedua program menunjukkan bahwa logika yang sama (penjumlahan lalu pembagian) dapat digunakan kembali hanya dengan menyesuaikan tipe parameter, mengurangi pengulangan kode dan risiko kesalahan (*bug*) akibat kode yang tidak konsisten.
