# <h1 align="center">Laporan Praktikum Modul 10 & 11 - Fungsi dan Prosedur</h1>
<p align="center">Raihan Ahmad Rahmadhani - 109092600025</p>

## Dasar Teori

### A. Konsep Fungsi (*Function*) dalam Bahasa Go
Fungsi (*function*) merupakan blok kode modular yang dirancang untuk melakukan tugas perhitungan atau pemrosesan data tertentu dan mengembalikan satu atau lebih nilai hasil (*return value*) kepada bagian program yang memanggilnya. Prinsip utama fungsi mengacu pada konsep matematis pemetaan $f(x) = y$, di mana masukan tertentu diproses untuk menghasilkan suatu keluaran spesifik.

Dalam bahasa Go, fungsi dideklarasikan dengan sintaks:
```go
func namaFungsi(parameter tipe) tipeKembalian {
    // instruksi pemrosesan
    return nilaiHasil
}
```
Karakteristik penting fungsi dalam Go:
1. **Nilai Kembalian Eksplisit**: Nilai hasil dikembalikan menggunakan kata kunci `return`. Nilai kembalian dapat langsung disematkan dalam ekspresi perhitungan atau ditampung ke dalam variabel penampung.
2. **Multiple Return Values**: Go mendukung pengembalian lebih dari satu nilai sekaligus dari satu fungsi tunggal.
3. **Murni (*Pure Function Potential*)**: Idealnya fungsi beroperasi atas masukan yang diberikan tanpa mengubah variabel di luar cakupannya (*side-effect free*).

### B. Konsep Prosedur (*Procedure*) dan Implementasinya di Go
Prosedur (*procedure*) adalah subprogram yang bertugas menjalankan suatu tindakan, aksi, pembaruan status data (*state modification*), atau menampilkan keluaran ke layar tanpa mengembalikan nilai hasil secara langsung melalui ekspresi pemanggil (*void return*).

Bahasa Go secara sintaksis tidak menyediakan kata kunci `procedure` khusus layaknya bahasa Pascal. Sebagai gantinya, prosedur diwujudkan melalui:
1. **Fungsi tanpa tipe kembalian**: Dideklarasikan tanpa tipe data setelah daftar parameter (`func namaProsedur(parameter)`).
2. **Pemanggilan sebagai instruksi mandiri**: Prosedur dipanggil langsung sebagai satu baris perintah berdiri sendiri, bukan sebagai bagian dari ekspresi aritmatika.

### C. Mekanisme Pengiriman Parameter: *Pass-by-Value* vs *Pass-by-Reference*
Dalam perancangan subprogram, mekanisme pengiriman argumen ke parameter menentukan pengaruh operasi internal subprogram terhadap data pemanggil:

1. **Parameter Masukan (*Input Parameter / By-Value*)**:
   - Nilai variabel disalin (*copy*) ke memori lokal subprogram.
   - Perubahan nilai parameter di dalam subprogram tidak memengaruhi nilai asli pada pemanggil.
   - Digunakan pada fungsi konvensional serta prosedur yang hanya membaca data (*read-only*).

2. **Parameter Keluaran (*Output Parameter / By-Reference*)**:
   - Digunakan untuk mengembalikan nilai komputasi dari prosedur ke pemanggil.
   - Nilai awal pada pemanggil belum bermakna (*Initial State* sembarang/kosong), lalu diisi oleh prosedur (*Final State* terdefinisi).
   - Di Go, mekanisme ini diimplementasikan menggunakan **pointer** (`*tipe`) dengan mengirimkan alamat memori (`&variabel`).

3. **Parameter Masukan/Keluaran (*Input-Output Parameter / By-Reference*)**:
   - Membawa data awal dari pemanggil untuk diolah, lalu hasil modifikasinya langsung memperbarui variabel asli di pemanggil.
   - Diimplementasikan di Go menggunakan pointer (`*tipe`). Operator dereferensi (`*`) digunakan di dalam badan subprogram untuk memodifikasi langsung nilai pada memori yang ditunjuk.

### D. Perbandingan Komparatif Fungsi dan Prosedur

| Aspek | Fungsi (*Function*) | Prosedur (*Procedure*) |
|---|---|---|
| **Nilai Kembalian** | Memiliki tipe kembalian formal (`return`). | Tidak memiliki tipe kembalian (*void*). |
| **Bentuk Pemanggilan** | Bagian dari ekspresi: `hasil := hitung(x)`. | Pernyataan berdiri sendiri: `cetak(x)`. |
| **Tujuan Utama** | Menghitung dan menghasilkan suatu nilai baru. | Menjalankan aksi, menampilkan output, atau mengubah *state*. |
| **Pengiriman Hasil** | Melalui kata kunci `return`. | Melalui parameter `out` atau `in/out` berbasis pointer (`*`). |

### E. Pemanfaatan *Struct* dan *Type Alias* Array dalam Subprogram
Ketika data yang diproses memiliki struktur majemuk (seperti data pemain dengan nama dan skor, atau gempa dengan nama, tahun, dan magnitudo), Go memanfaatkan `struct` untuk mengelompokkan variabel terkait ke dalam satu kesatuan logis.

Untuk menangani kumpulan data majemuk, didefinisikan tipe data array melalui *type alias*:
```go
const NMAX = 1000
type Mahasiswa struct {
    nama string
    nilai int
}
type TabMhs = [NMAX]Mahasiswa
```
Dengan mendefinisikan *type alias*, array tersebut dapat dilewatkan sebagai parameter fungsi maupun prosedur secara rapi, jelas, dan aman (*type-safe*).

---

## Guided

### 1. player.go (Pertemuan 10 - Fungsi)

```go
package main

import "fmt"

const NMAX = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer = [NMAX]player

// function rataRata(T : tabPlayer, n : integer) -> real
// mengembalikan rata-rata skor dari n player
func rataRata(T tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].score
	}
	return float64(total) / float64(n)
}

// function idxTertinggi(T : tabPlayer, n : integer) -> integer
// mengembalikan indeks player dengan skor tertinggi;
// jika ada skor sama, ambil yang pertama muncul
func idxTertinggi(T tabPlayer, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].score > T[idxMax].score {
			idxMax = i
		}
	}
	return idxMax
}

// function hitungDiAtas(T : tabPlayer, n : integer, batas : real) -> integer
// mengembalikan banyaknya player dengan skor lebih dari batas
func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var T tabPlayer

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	rata := rataRata(T, n)
	terbaik := idxTertinggi(T, n)
	diAtas := hitungDiAtas(T, n, rata)

	fmt.Printf("Rata-rata skor: %.2f\n", rata)
	fmt.Printf("Player terbaik: %s (%d)\n", T[terbaik].nickname, T[terbaik].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", diAtas)
}
```

##### Output
![Screenshot Output Guided 1](guided/output.png)

#### Deskripsi
Program `player.go` mengimplementasikan pengolahan data turnamen game menggunakan konsep modularitas berbasis fungsi dengan nilai kembalian (*return value*).

1. **Struktur Data dan Tipe**:
   - Didefinisikan konstanta `NMAX = 999`.
   - `struct player` memuat atribut `nickname` bertipe string dan `score` bertipe integer.
   - `tabPlayer` merupakan *type alias* untuk array berukuran `NMAX` elemen `player`.

2. **Fungsi `rataRata(T tabPlayer, n int) float64`**:
   - Menghitung akumulasi seluruh skor pemain dari indeks `0` hingga `n-1`.
   - Mengembalikan nilai rata-rata bertipe `float64` hasil pembagian total skor dengan jumlah pemain `n`.

3. **Fungsi `idxTertinggi(T tabPlayer, n int) int`**:
   - Mengasumsikan indeks pemain terbaik awal berada pada posisi `0` (`idxMax := 0`).
   - Melakukan iterasi dari indeks `1` sampai `n-1`. Jika ditemukan skor yang lebih besar secara tegas (`T[i].score > T[idxMax].score`), maka `idxMax` diperbarui.
   - Penggunaan operator pembanding `>` memastikan bahwa jika terdapat pemain dengan skor tertinggi yang bernilai kembar, pemain yang pertama kali muncul yang terpilih.

4. **Fungsi `hitungDiAtas(T tabPlayer, n int, batas float64) int`**:
   - Melakukan pencacahan berapa banyak pemain yang memiliki skor lebih besar dari parameter `batas` (rata-rata skor).
   - Mengembalikan bilangan bulat berupa total pemain yang memenuhi kriteria.

5. **Fungsi `main`**:
   - Membaca jumlah pemain `n` dan data masing-masing pemain ke dalam array `T`.
   - Memanggil ketiga fungsi di atas secara berurutan dan menampung nilai kembaliannya.
   - Mencetak hasil rata-rata skor dengan format 2 angka di belakang koma (`%.2f`), nama dan skor pemain terbaik, serta banyaknya pemain yang berada di atas rata-rata.

---

### 2. sapi.go (Pertemuan 11 - Prosedur)

```go
package main

import "fmt"

const NMAX = 10000

type Arr = [NMAX]int

// procedure bacaData(in/out T : Arr, in n : integer)
// IS : n terdefinisi, T belum berisi data
// FS : T[0..n-1] berisi n nomor sapi hasil input
func bacaData(T *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

// procedure geserKanan(in/out T : Arr, in n : integer)
// IS : T[0..n-1] berisi susunan sapi saat ini
// FS : setiap sapi bergeser satu kandang ke kanan dan
// sapi di kandang terakhir pindah ke kandang pertama
func geserKanan(T *Arr, n int) {
	temp := T[n-1]
	for i := n - 1; i > 0; i-- {
		T[i] = T[i-1]
	}
	T[0] = temp
}

// procedure cetak(in T : Arr, in n : integer)
// IS : T[0..n-1] terdefinisi
// FS : isi T tercetak dalam satu baris dipisahkan spasi
func cetak(T Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var n, k int
	var T Arr

	fmt.Scan(&n, &k)
	bacaData(&T, n)

	// Nilai K dapat mencapai 10^9. Pergeseran sebanyak N kali akan mengembalikan
	// posisi sapi ke susunan awal. Oleh karena itu, cukup dilakukan pergeseran
	// sebanyak (K % N) kali agar program berjalan efisien (O(N) bukannya O(K*N)).
	rotasi := k % n
	for i := 0; i < rotasi; i++ {
		geserKanan(&T, n)
	}

	cetak(T, n)
}
```

##### Output
![Screenshot Output Guided 2](guided/output2.png)

#### Deskripsi
Program `sapi.go` memecahkan masalah rotasi sirkular posisi sapi di $N$ kandang setelah $K$ hari dengan memanfaatkan prosedur murni (fungsi tanpa nilai kembalian dengan parameter pointer).

1. **Prosedur `bacaData(T *Arr, n int)`**:
   - Memiliki tipe parameter `in/out T : Arr`.
   - Menggunakan pointer `*Arr` sehingga masukan yang dibaca langsung tersimpan pada variabel array asli di `main`.

2. **Prosedur `geserKanan(T *Arr, n int)`**:
   - Menggeser seluruh elemen array satu posisi ke kanan.
   - Nilai elemen terakhir `T[n-1]` disimpan dalam variabel sementara `temp`.
   - Perulangan mundur memindahkan elemen `T[i-1]` ke `T[i]`, dan akhirnya `temp` diletakkan pada posisi pertama `T[0]`.

3. **Prosedur `cetak(T Arr, n int)`**:
   - Parameter `T Arr` berjenis `in` (*pass-by-value*) karena hanya membaca dan menampilkan elemen tanpa mengubah susunan data.

4. **Optimasi Algoritmik Modulus**:
   - Batas nilai $K$ mencapai $10^9$. Jika pergeseran dilakukan $K$ kali putaran fisik, program akan mengalami *Time Limit Exceeded* (TLE).
   - Mengingat susunan array akan berulang setiap $N$ kali pergeseran sirkular, jumlah rotasi efektif dihitung dengan `rotasi := k % n`. Dengan demikian, prosedur `geserKanan` paling banyak hanya dipanggil $N-1$ kali, memberikan performa optimal $O(N)$.

---

## Unguided

### 1. gempa.go (Pertemuan 10 - Fungsi)

```go
package main

import "fmt"

const NMAX = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa = [NMAX]gempa

// function kategori(m : real) -> string
func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m < 6.0 {
		return "Sedang"
	} else if m < 7.0 {
		return "Kuat"
	} else {
		return "Besar"
	}
}

// function rataMagnitudo(T : tabGempa, n : integer) -> real
func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}
	return total / float64(n)
}

// function idxTerkuat(T : tabGempa, n : integer) -> integer
// jika ada lebih dari satu gempa dengan magnitudo terbesar,
// gempa terkuat adalah yang pertama muncul
func idxTerkuat(T tabGempa, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[idxMax].magnitudo {
			idxMax = i
		}
	}
	return idxMax
}

// function hitungKategori(T : tabGempa, n : integer, k : string) -> integer
// fungsi hitungKategori wajib memanggil fungsi kategori di dalamnya
func hitungKategori(T tabGempa, n int, k string) int {
	count := 0
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var T tabGempa

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n", T[i].nama, T[i].tahun, T[i].magnitudo, kategori(T[i].magnitudo))
	}

	terkuat := idxTerkuat(T, n)
	fmt.Printf("Rata-rata magnitudo: %.2f\n", rataMagnitudo(T, n))
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", T[terkuat].nama, T[terkuat].tahun, T[terkuat].magnitudo)
	fmt.Printf("Jumlah gempa kategori Besar: %d\n", hitungKategori(T, n, "Besar"))
}
```

##### Output
![Screenshot Output Unguided 1](unguided/output.png)

#### Deskripsi
Program `gempa.go` dirancang untuk mengolah dan menganalisis sekumpulan catatan data seismologi (gempa bumi) melalui pemisahan logika ke dalam beberapa fungsi terspesialisasi.

1. **Struktur Data `gempa`**:
   - Memiliki tiga atribut: `nama` (wilayah, string), `tahun` (integer), dan `magnitudo` (skala gempa, float64).
   - Tipe alias `tabGempa` menampung hingga `NMAX = 999` rekaman gempa.

2. **Fungsi `kategori(m float64) string`**:
   - Mengklasifikasikan skala magnitudo gempa:
     - $m < 4.0 \rightarrow$ `"Minor"`
     - $4.0 \le m < 6.0 \rightarrow$ `"Sedang"`
     - $6.0 \le m < 7.0 \rightarrow$ `"Kuat"`
     - $m \ge 7.0 \rightarrow$ `"Besar"`

3. **Fungsi `rataMagnitudo(T tabGempa, n int) float64`**:
   - Menghitung nilai rerata magnitudo dari $n$ data kejadian gempa.

4. **Fungsi `idxTerkuat(T tabGempa, n int) int`**:
   - Mencari indeks kejadian gempa berkekuatan terbesar. Apabila terdapat dua gempa dengan magnitudo puncak yang sama, fungsi mempertahankan gempa yang pertama kali dimasukkan (menggunakan pembanding `>`).

5. **Fungsi `hitungKategori(T tabGempa, n int, k string) int`**:
   - Menghitung frekuensi kemunculan kategori tertentu (misalnya `"Besar"`). Sesuai ketentuan spesifikasi, fungsi ini **wajib memanggil fungsi `kategori`** di dalam perulangannya untuk menentukan label kategori setiap elemen secara konsisten.

6. **Alur Fungsi `main`**:
   - Membaca data masukan, kemudian mencetak daftar rekaman gempa beserta label kategorinya.
   - Menampilkan rata-rata magnitudo (format `%.2f`), informasi lengkap gempa terkuat, serta total kejadian gempa berkategori `"Besar"`.

---

### 2. nim.go (Pertemuan 11 - Prosedur)

```go
package main

import "fmt"

const NMAX = 2002

type tabNIM = [NMAX]int

// procedure cariNIM(in T : tabNIM, in n, x : integer, out idx : integer)
// IS : T[0..n-1] dan x terdefinisi
// FS : idx berisi indeks NIM x pada T, atau -1 jika tidak ditemukan
func cariNIM(T tabNIM, n, x int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i] == x {
			*idx = i
			return
		}
	}
}

// procedure ubahNIM(in/out T : tabNIM, in idx, baru : integer)
// IS : T terdefinisi dan idx adalah indeks yang valid
// FS : T[idx] berisi NIM baru
func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}

// procedure cetak(in T : tabNIM, in n : integer)
// IS : T[0..n-1] terdefinisi
// FS : setiap NIM pada T tercetak satu per baris
func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var n int
	var T tabNIM
	var cari, baru, idx int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	fmt.Scan(&cari)
	cariNIM(T, n, cari, &idx)

	if idx != -1 {
		fmt.Scan(&baru)
		ubahNIM(&T, idx, baru)
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", cari, idx+1)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}
```

##### Output
![Screenshot Output Unguided 2](unguided/output2.png)

#### Deskripsi
Program `nim.go` mensimulasikan mekanisme pencarian dan pembaruan data Nomor Induk Mahasiswa (NIM) pada sistem akademik menggunakan konsep prosedur dan parameter pointer.

1. **Prosedur `cariNIM(T tabNIM, n, x int, idx *int)`**:
   - Memiliki parameter output `idx *int`.
   - Menginisialisasi `*idx = -1`. Melakukan pencarian sekuensial (*linear search*); jika ditemukan $T[i] = x$, alamat yang ditunjuk `idx` diisi dengan indeks $i$ (`*idx = i`) dan prosedur langsung berhenti.

2. **Prosedur `ubahNIM(T *tabNIM, idx, baru int)`**:
   - Memiliki parameter input/output `T *tabNIM`.
   - Memperbarui elemen pada indeks `idx` dengan nilai `baru` secara langsung pada alamat memori array pemanggil (`T[idx] = baru`).

3. **Prosedur `cetak(T tabNIM, n int)`**:
   - Menampilkan seluruh NIM yang tersimpan pada array satu per satu per baris.

4. **Alur Logika Fungsi `main`**:
   - Membaca $N$ data NIM, lalu membaca NIM target yang ingin dicari (`cari`).
   - Memanggil `cariNIM(T, n, cari, &idx)`.
   - Jika `idx != -1`, program membaca NIM baru (`baru`), melakukan pembaruan data via `ubahNIM`, mencetak posisi mahasiswa ke-`idx+1` (karena indeks array berbasis 0 sedangkan nomor urut berbasis 1), serta mencetak daftar NIM hasil perubahan.
   - Jika `idx == -1`, program mencetak `"NIM tidak ditemukan dalam sistem"` tanpa meminta input NIM baru.

---

## Kesimpulan
Dari praktikum gabungan Modul 10 (Fungsi) dan Modul 11 (Prosedur) dalam bahasa pemrograman Go, dapat ditarik beberapa kesimpulan penting:

1. **Pemilihan Subprogram Berdasarkan Kebutuhan**:
   - **Fungsi** sangat ideal untuk subprogram yang bertugas melakukan kalkulasi murni atau pemetaan data untuk menghasilkan sebuah nilai spesifik yang langsung digunakan dalam ekspresi aritmatika atau percabangan (seperti perhitungan rata-rata, penentuan kategori, atau pencarian nilai ekstrem).
   - **Prosedur** sangat tepat digunakan untuk menjalankan aksi atau urutan operasi yang mengubah keadaan (*state*) program, melakukan manipulasi susunan elemen array, atau berinteraksi dengan masukan/keluaran (I/O).

2. **Peran Pointer dalam Mewujudkan Prosedur di Go**:
   - Ketiadaan kata kunci `procedure` di Go digantikan secara elegan menggunakan fungsi tanpa nilai kembalian (*void*) dipadukan dengan parameter pointer (`*tipe`).
   - Pointer memungkinkan subprogram mengubah nilai variabel asli pada fungsi pemanggil (*pass-by-reference*), baik untuk parameter keluaran (`out`) maupun parameter pembaruan (`in/out`).

3. **Komposisi Subprogram (*Function Calling Function*)**:
   - Fungsi dapat memanggil fungsi lain di dalamnya (seperti `hitungKategori` yang memanggil `kategori`), menciptakan rantai abstraksi kode yang kuat, mematuhi prinsip *Single Responsibility*, dan mencegah duplikasi aturan logika.

4. **Kombinasi *Type Alias* dan *Struct***:
   - Penggunaan `struct` memudahkan pengorganisasian entitas dunia nyata ke dalam pemrograman terstruktur, sementara *type alias* pada array mempermudah perlakuan kumpulan data sebagai satu tipe parameter yang seragam saat dilewatkan ke dalam subprogram.

5. **Efisiensi Algoritma pada Struktur Modular**:
   - Modularitas mempermudah penerapan optimasi terisolasi, seperti penggunaan operasi modulus $K \pmod N$ pada rotasi sirkular array (`sapi.go`) yang mampu memangkas waktu eksekusi dari $O(K \times N)$ menjadi $O(N)$ tanpa merusak struktur utama program.
