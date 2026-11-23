# <h1 align="center">Tugas Pendahuluan Modul 10 & 11 - Fungsi dan Prosedur</h1>
<p align="center">Raihan Ahmad Rahmadhani - 109092600025</p>

---

## Dasar Teori

### A. Konsep Fungsi dan Penggunaannya dalam Bahasa Go

#### Deskripsi
Fungsi (*function*) adalah blok kode terstruktur yang dirancang untuk melakukan tugas perhitungan atau pemrosesan tertentu dan mengembalikan suatu nilai hasil (*return value*) ke bagian program yang memanggilnya. Fungsi merupakan fondasi dalam dekomposisi modularitas yang memecah program kompleks menjadi sub-subprogram kecil yang mandiri, mudah dirawat, dan dapat digunakan berulang kali (*reusable*).

Dalam bahasa Go, fungsi didefinisikan menggunakan kata kunci `func` dengan sintaks:
```go
func namaFungsi(parameter tipe) tipePengembalian {
    // instruksi
    return nilai
}
```

**Komponen utama fungsi dalam Go:**
1. **Nama fungsi** — Identifier pemanggil fungsi. Menggunakan huruf kapital diawal jika *exported* (dapat diakses dari luar *package*), atau huruf kecil jika hanya untuk *package* lokal.
2. **Parameter** — Masukan yang diterima fungsi beserta deklarasi tipe datanya secara eksplisit.
3. **Tipe kembalian (*return type*)** — Tipe data dari nilai yang dihasilkan fungsi. Go mendukung pengembalian nilai tunggal maupun jamak (*multiple return values*).
4. **Pernyataan `return`** — Mengembalikan nilai hasil kalkulasi ke pemanggil fungsi.

---

### B. Perbedaan Prosedur dan Fungsi serta Penggunaan Parameter Pointer pada Go

#### Deskripsi
Dalam paradigma pemrograman prosedural, terdapat dua konstruksi utama subprogram: **fungsi** (*function*) dan **prosedur** (*procedure*):

1. **Fungsi (*Function*):**
   - Subprogram yang **mengembalikan nilai hasil** (*return value*) melalui `return`.
   - Dipanggil sebagai bagian dari suatu ekspresi, misalnya `hasil := hitungLuas(p, l)`.
   - Bertujuan menghitung nilai murni tanpa menimbulkan efek samping pada variabel pemanggil.

2. **Prosedur (*Procedure*):**
   - Subprogram yang **tidak memiliki tipe kembalian** (*void*).
   - Bekerja dengan menjalankan urutan aksi atau memodifikasi variabel di pemanggil melalui mekanisme ***pass-by-reference*** menggunakan pointer.
   - Dipanggil sebagai pernyataan mandiri, misalnya `tukar(&a, &b)`.

3. **Implementasi Prosedur di Go Menggunakan Parameter Pointer:**
   - Go tidak memiliki kata kunci khusus `procedure`. Prosedur diimplementasikan sebagai fungsi tanpa tipe kembalian di sebelah kanan tanda kurung parameter.
   - Parameter keluaran (*out*) atau pembaruan (*in/out*) dikirimkan sebagai tipe **pointer** (`*tipe`). Saat dipanggil, argumen dilewatkan dengan menyertakan alamat memorinya melalui operator `&` (*address-of*).
   - Di dalam badan prosedur, operator dereferensi `*` digunakan untuk memodifikasi langsung nilai pada memori variabel pemanggil.

---

## Pembahasan Soal Tugas Pendahuluan

### 1. TP Soal 1: raport.go

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

![Screenshot Output TP Soal 1](TP/soal%201/output.png)

#### Deskripsi
Program `raport.go` mencetak laporan nilai rapor mata kuliah mahasiswa dengan memanfaatkan fungsi modular untuk konversi nilai numerik ke huruf serta kalkulasi rerata nilai.

1. Konstanta `NMAX = 999` membatasi kapasitas penyimpanan mata kuliah.
2. `struct rapot` mengelompokkan data nama mata kuliah (`matkul`, string) dan perolehan skornya (`nilai`, int). Tipe `tabRapot` adalah alias array berukuran `NMAX` bertipe `rapot`.
3. Fungsi `nilaiHuruf(x int) string` menerima nilai numerik dan mengembalikan grade huruf bersyarat: $>80 \rightarrow \text{"A"}$, $>70 \rightarrow \text{"AB"}$, $>65 \rightarrow \text{"B"}$, $>60 \rightarrow \text{"BC"}$, $>50 \rightarrow \text{"C"}$, $>40 \rightarrow \text{"D"}$, dan selain itu $\rightarrow \text{"E"}$.
4. Fungsi `rataRata(T tabRapot, n int) float64` mengakumulasikan nilai dari $n$ data pertama array $T$ lalu membaginya dengan $n$.
5. Pada `main`, setelah membaca $n$ dan mengisi data rapor, program menampilkan setiap mata kuliah, nilai angka, nilai huruf hasil evaluasi `nilaiHuruf`, serta rata-rata akhir dengan format dua desimal (`%.2f`).

---

### 2. TP Soal 2: inputnilai.go

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

![Screenshot Output TP Soal 2](TP/soal%202/output.png)

#### Deskripsi
Program `inputnilai.go` mengelola pencatatan data akademik mahasiswa yang dilengkapi fungsi pencegah duplikasi data NIM dan penghitungan nilai rata-rata kelas.

1. Didefinisikan `struct student` yang memuat field `nim` (string), `name` (string), dan `grade` (float64). `students` merupakan alias array bertipe `student` berkapasitas `NMAX = 100`.
2. Fungsi `cariNIM(T students, n int, nim string) int` menelusuri array $T$ secara sekuensial. Jika NIM ditemukan, indeks posisinya dikembalikan; jika tidak ditemukan, fungsi mengembalikan `-1`.
3. Fungsi `rataRata(T students, n int) float64` menjumlahkan field `grade` dari $n$ mahasiswa lalu membaginya dengan $n$.
4. Pada fungsi `main`, masukan dibaca secara berulang dalam *infinite loop* yang berhenti ketika menerima input sentinel `"STOP"`.
5. Setiap pasang data baru diverifikasi menggunakan pemanggilan `cariNIM`. Data hanya dimasukkan ke array jika NIM belum pernah terdaftar (`cariNIM == -1`).
6. Setelah pembacaan selesai, program menampilkan daftar mahasiswa yang berhasil tersimpan serta mencetak nilai rata-rata kelas.

---

### 3. TP Soal 3: sekolah.go

```go
package main

import "fmt"

const NMAX = 1001

type tabMurid = [NMAX]string

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false
	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1
			return
		}
	}
}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			*jumlah++
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabMurid
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	var nama string
	fmt.Scan(&nama)

	var posisi int
	var ketemu bool
	cariMurid(T, n, nama, &posisi, &ketemu)

	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}

	var jumlah int
	hitungHuruf(T, n, nama[0], &jumlah)
	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", nama[0], jumlah)
}
```

##### Output

![Screenshot Output TP Soal 3](TP/soal%203/output.png)

#### Deskripsi
Program `sekolah.go` mengelola absensi nama murid menggunakan prosedur murni dengan parameter keluaran berbasis pointer (*pass-by-reference*).

1. Tipe `tabMurid` merupakan alias array string berkapasitas `NMAX = 1001`.
2. **Prosedur `cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool)`**:
   - Memiliki dua parameter keluaran berbasis pointer: `*posisi` dan `*ketemu`.
   - Menginisialisasi `*ketemu = false`. Melakukan pencarian sekuensial pada array $T$. Jika nama cocok, `*ketemu` diubah menjadi `true`, nomor urut absen diisi `*posisi = i + 1`, lalu prosedur langsung keluar (`return`).
3. **Prosedur `hitungHuruf(T tabMurid, n int, h byte, jumlah *int)`**:
   - Menghitung berapa banyak murid yang namanya diawali oleh karakter huruf tertentu `h`.
   - Setiap kecocokan `T[i][0] == h` menginkremen nilai yang ditunjuk oleh pointer `*jumlah`.
4. Pada `main`, kedua prosedur dipanggil dengan melewatkan alamat memori variabel lokal (`&posisi`, `&ketemu`, `&jumlah`). Perubahan data di dalam prosedur secara langsung memengaruhi nilai variabel di fungsi `main`.

---

### 4. TP Soal 4: penduduk.go

```go
package main

import "fmt"

const NMAX = 1000

type penduduk struct {
	nama  string
	tahun int
	kota  string
}

type tabPenduduk = [NMAX]penduduk

func cariTahun(T tabPenduduk, n, th int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i].tahun == th {
			*idx = i
			return
		}
	}
}

func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
	*tua = T[0]
	*muda = T[0]
	for i := 1; i < n; i++ {
		if T[i].tahun < tua.tahun {
			*tua = T[i]
		}
		if T[i].tahun > muda.tahun {
			*muda = T[i]
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabPenduduk
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
	}

	var th int
	fmt.Scan(&th)

	var idx int
	cariTahun(T, n, th, &idx)
	if idx != -1 {
		fmt.Printf("%s %d %s ditemukan di index ke-%d\n", T[idx].nama, T[idx].tahun, T[idx].kota, idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	var tua, muda penduduk
	tertuaTermuda(T, n, &tua, &muda)
	fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}
```

##### Output

![Screenshot Output TP Soal 4](TP/soal%204/output.png)

#### Deskripsi
Program `penduduk.go` mengelola pencatatan sensus kependudukan menggunakan struct majemuk dan prosedur yang menerima parameter pointer ke struct.

1. `struct penduduk` mengelompokkan atribut `nama` (string), `tahun` (int), dan `kota` (string). `tabPenduduk` merupakan alias array berkapasitas `NMAX = 1000`.
2. **Prosedur `cariTahun(T tabPenduduk, n, th int, idx *int)`**:
   - Mencari penduduk berdasarkan tahun lahir `th`.
   - Menuliskan indeks yang ditemukan ke parameter pointer `*idx`, atau tetap `-1` jika tahun tidak ditemukan dalam daftar.
3. **Prosedur `tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk)`**:
   - Menerima dua parameter pointer ke struct: `tua *penduduk` dan `muda *penduduk`.
   - Menginisialisasi `*tua` dan `*muda` dengan elemen pertama `T[0]`. Melalui iterasi, membandingkan nilai `tahun`: tahun lebih kecil mengindikasikan penduduk lebih tua, sedangkan tahun lebih besar mengindikasikan penduduk lebih muda.
   - Karena parameter berupa pointer ke struct, seluruh field data (nama, tahun, kota) langsung termodifikasi pada variabel pemanggil tanpa overhead salinan data.
4. Pada `main`, program menampilkan data penduduk hasil pencarian tahun serta mencetak nama dan tahun lahir penduduk tertua dan termuda.

---

## Kesimpulan

1. **Pemilihan Fungsi vs Prosedur**: Fungsi digunakan ketika subprogram bertujuan menghasilkan satu nilai kalkulasi yang langsung digunakan dalam evaluasi ekspresi (seperti `rataRata` dan `nilaiHuruf`), sedangkan prosedur digunakan ketika subprogram menjalankan aksi, memodifikasi struktur data, atau menghasilkan beberapa keluaran sekaligus melalui parameter pointer (seperti `cariMurid`, `tertuaTermuda`).
2. **Implementasi Prosedur di Go**: Bahasa Go tidak memiliki kata kunci `procedure`. Prosedur di Go direalisasikan melalui fungsi bertipe kembalian kosong (*void*) yang memanfaatkan parameter **pointer** (`*tipe`) dan pengiriman alamat memori (`&variabel`) untuk mekanisme *pass-by-reference*.
3. **Efisiensi Penggunaan Pointer**: Pada tipe data berukuran besar seperti array dan struct, melewatkan pointer ke dalam subprogram menghemat penggunaan memori komputasi karena program tidak perlu menduplikasi seluruh isi elemen saat pemanggilan subprogram berlangsung.
