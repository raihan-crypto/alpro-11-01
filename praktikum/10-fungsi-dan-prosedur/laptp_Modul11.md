# <h1 align="center">Tugas Pendahuluan Modul 11 - Prosedur</h1>
<p align="center">Raihan Ahmad Rahmadhani - 109092600025</p>

### 1. Perbedaan Prosedur dan Fungsi serta Penggunaan Parameter Pointer pada Go

#### Deskripsi
Dalam paradigma pemrograman prosedural, terdapat dua konstruksi utama untuk membungkus dan mengorganisasikan blok instruksi yang dapat dipanggil ulang: **fungsi** (*function*) dan **prosedur** (*procedure*). Meskipun keduanya sama-sama memungkinkan dekomposisi program menjadi unit-unit yang lebih kecil dan terstruktur, keduanya memiliki perbedaan mendasar dalam hal nilai kembalian dan cara berinteraksi dengan data di luar cakupannya.

Berikut adalah penjelasan perbedaan antara prosedur dan fungsi, serta cara Go mengimplementasikan prosedur menggunakan parameter pointer:

1. **Fungsi (*Function*):**
   - Sebuah subprogram yang **mengembalikan satu nilai hasil** (*return value*) kepada pemanggil melalui pernyataan `return`.
   - Nilai kembalian menjadi bagian dari ekspresi; pemanggil dapat langsung menggunakan hasil fungsi dalam operasi lain, misalnya `hasil := hitungLuas(5, 10)`.
   - Cocok digunakan ketika sebuah subprogram bertujuan menghitung dan menghasilkan satu nilai tunggal yang langsung dipakai.
   - **Sintaks pada Go:**
     ```go
     func namaFungsi(param tipe) tipekembalian {
         // instruksi
         return nilaiHasil
     }
     ```

2. **Prosedur (*Procedure*):**
   - Sebuah subprogram yang **tidak mengembalikan nilai** kepada pemanggil (bertipe `void` dalam istilah C, atau tidak ada tipe kembalian).
   - Prosedur bekerja dengan cara memodifikasi variabel yang ada di luar cakupannya, biasanya melalui mekanisme **pass-by-reference** menggunakan pointer, sehingga efek dari prosedur tampak pada variabel yang dioper.
   - Cocok digunakan ketika sebuah subprogram bertujuan melakukan tindakan atau memodifikasi beberapa variabel sekaligus tanpa perlu mengembalikan nilai.
   - **Sintaks pada Go (menggunakan pointer sebagai parameter):**
     ```go
     func namaProsedur(param tipe, hasil *tipe) {
         // instruksi
         *hasil = nilaiHasil
     }
     ```

3. **Implementasi Prosedur di Go dengan Parameter Pointer:**
   - Bahasa Go tidak memiliki kata kunci khusus `procedure`. Prosedur diimplementasikan sebagai fungsi bertipe kembalian kosong (tanpa tipe di kanan nama fungsi), yang menerima parameter bertipe **pointer** (`*tipe`) untuk menyimpan hasil perhitungan.
   - Ketika sebuah variabel dioper sebagai pointer dengan operator `&` (address-of), fungsi yang menerima pointer tersebut dapat langsung mengubah nilai variabel asli di luar fungsi menggunakan operator dereference `*`.
   - Mekanisme ini dikenal sebagai ***pass-by-reference***: yang dikirimkan bukan nilainya, melainkan **alamat memori** variabel tersebut, sehingga perubahan di dalam prosedur langsung memengaruhi variabel aslinya.
   - **Contoh penggunaan:**
     ```go
     func hitungKuadrat(x int, hasil *int) {
         *hasil = x * x
     }

     func main() {
         var kuadrat int
         hitungKuadrat(5, &kuadrat)
         fmt.Println(kuadrat) // Output: 25
     }
     ```

---

#### Dua Perbedaan Prosedur pada Go Dibandingkan dengan Bahasa Lain (Pascal / C++)

1. **Tidak Ada Kata Kunci Khusus `procedure`:**
   - **Pada Pascal:** Bahasa Pascal secara tegas membedakan subprogram yang mengembalikan nilai (didefinisikan dengan kata kunci `function`) dan yang tidak mengembalikan nilai (didefinisikan dengan kata kunci `procedure`). Pembedaan ini bersifat sintaksis dan wajib.
   - **Pada Go:** Go tidak mengenal kata kunci `procedure`. Semua subprogram — baik yang mengembalikan nilai maupun tidak — dideklarasikan menggunakan kata kunci `func`. Prosedur diimplementasikan hanya dengan tidak menuliskan tipe kembalian, sehingga fungsi bertipe kembalian kosong secara konvensi berperan sebagai prosedur. Parameter pointer kemudian digunakan untuk meneruskan hasil keluar.

2. **Multiple Return Values sebagai Alternatif Modern:**
   - **Pada C++ / Java:** Ketika sebuah subprogram perlu menghasilkan lebih dari satu keluaran, programmer harus menggunakan pointer, referensi (`&`), atau objek pembungkus (*wrapper class*). Ini membuat kode lebih rumit dan berpotensi menimbulkan kesalahan pengelolaan memori.
   - **Pada Go:** Selain pendekatan prosedur dengan pointer, Go mendukung ***multiple return values*** secara native, yaitu sebuah fungsi dapat langsung mengembalikan beberapa nilai sekaligus dengan sintaks `return nilai1, nilai2`. Ini menjadi fitur idiomatis Go yang lebih bersih dan aman dibandingkan menggunakan pointer untuk output berganda.

---

### 2. sekolah.go

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

```text
Masukan:
5
Ani Budi Cahya Ani Dani
Cahya

Keluaran:
Murid terdaftar dan berada di urutan absen ke-3
Jumlah murid dengan huruf awal C: 1
```

#### Deskripsi
Program `sekolah.go` ini mengelola daftar nama murid dalam sebuah array dan menyediakan dua prosedur utama: mencari posisi absen seorang murid berdasarkan nama, serta menghitung berapa murid yang namanya diawali dengan huruf yang sama dengan nama yang dicari.

Cara kerja program:
1. Program mendefinisikan tipe `tabMurid` sebagai alias array `[NMAX]string` dengan kapasitas maksimum `NMAX = 1001` elemen, lalu membaca jumlah murid `n` dan seluruh nama dari terminal.
2. **Prosedur `cariMurid`** menerima lima parameter: array `T`, jumlah elemen `n`, nama yang dicari `nama`, serta dua parameter pointer `*posisi` dan `*ketemu` sebagai keluaran.
   - Prosedur ini menelusuri seluruh elemen array satu per satu dalam perulangan `for`.
   - Jika elemen yang diperiksa cocok dengan `nama`, nilai `*ketemu` di-set `true`, nilai `*posisi` di-isi dengan nomor urut absen (`i + 1`, karena indeks mulai dari 0), dan prosedur langsung keluar dengan `return`.
   - Jika seluruh iterasi selesai tanpa kecocokan, `*ketemu` tetap bernilai `false` dan `*posisi` tidak diubah.
3. **Prosedur `hitungHuruf`** menerima empat parameter: array `T`, jumlah elemen `n`, huruf awal `h` bertipe `byte`, dan pointer `*jumlah` sebagai keluaran.
   - Prosedur menelusuri seluruh elemen array; setiap kali karakter pertama suatu nama (`T[i][0]`) sama dengan `h`, variabel yang ditunjuk `*jumlah` ditambah satu.
   - Pengecekan `len(T[i]) > 0` diperlukan untuk menghindari akses indeks pada string kosong.
4. Di fungsi `main`, kedua prosedur dipanggil dengan meneruskan alamat variabel lokal menggunakan operator `&`, membuktikan mekanisme *pass-by-reference*: perubahan di dalam prosedur langsung tercermin pada variabel di `main`.
5. Hasil akhirnya, program mencetak apakah murid ditemukan beserta nomor absennya, dilanjutkan dengan jumlah murid yang berbagi huruf awal yang sama dengan nama yang dicari.

---

### 3. penduduk.go

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

```text
Masukan:
4
Budi 1950 Jakarta
Sari 1990 Bandung
Andi 1975 Surabaya
Dewi 2001 Medan
1975

Keluaran:
Andi 1975 Surabaya ditemukan di index ke-2
Penduduk tertua: Budi (1950)
Penduduk termuda: Dewi (2001)
```

#### Deskripsi
Program `penduduk.go` ini mengelola data kependudukan dalam sebuah array bertipe `struct` dan menyediakan dua prosedur: mencari penduduk berdasarkan tahun lahir, serta menentukan penduduk tertua dan termuda dari seluruh data yang ada.

Cara kerja program:
1. Program mendefinisikan tipe `penduduk` sebagai `struct` dengan tiga field: `nama` (string), `tahun` (int), dan `kota` (string). Tipe `tabPenduduk` adalah alias array `[NMAX]penduduk` dengan kapasitas maksimum `NMAX = 1000` elemen.
2. Program membaca jumlah data `n` kemudian mengisi array `T` dengan data setiap penduduk menggunakan `fmt.Scan` yang membaca `nama`, `tahun`, dan `kota` sekaligus per iterasi. Setelah itu, dibaca satu lagi nilai `th` sebagai tahun yang akan dicari.
3. **Prosedur `cariTahun`** menerima array `T`, jumlah elemen `n`, tahun target `th`, dan pointer keluaran `*idx`.
   - Nilai awal `*idx` disetel `-1` sebagai penanda "tidak ditemukan".
   - Prosedur menelusuri array; saat menemukan penduduk dengan `T[i].tahun == th`, nilai `*idx` diperbarui menjadi indeks `i` dan prosedur langsung keluar.
   - Karena return lebih awal dilakukan begitu data pertama yang cocok ditemukan, prosedur ini bersifat *first-match search*.
4. **Prosedur `tertuaTermuda`** menerima array `T`, jumlah elemen `n`, dan dua pointer keluaran `*tua` dan `*muda` bertipe `*penduduk`.
   - Keduanya diinisialisasi dengan data elemen pertama `T[0]` sebagai nilai pembanding awal.
   - Perulangan `for` mulai dari indeks `1` membandingkan setiap elemen:  penduduk dengan tahun lahir lebih kecil (lebih tua) menggantikan `*tua`, dan penduduk dengan tahun lahir lebih besar (lebih muda) menggantikan `*muda`.
   - Karena parameter bertipe pointer ke struct (`*penduduk`), seluruh data struct (nama, tahun, kota) langsung tersalin ke variabel di `main` ketika prosedur mengubah `*tua` atau `*muda`.
5. Hasil akhirnya, program mencetak data penduduk yang tahunnya cocok beserta indeksnya, kemudian mencetak nama dan tahun lahir penduduk tertua serta termuda.

---

## Kesimpulan

Dari pengerjaan Tugas Pendahuluan Modul 11 mengenai prosedur dalam bahasa pemrograman Go ini, beberapa poin penting yang dapat disimpulkan antara lain:

1. **Pemahaman Prosedur sebagai Subprogram Tanpa Nilai Kembalian:**
   - Prosedur adalah subprogram yang bertujuan melakukan tindakan atau memodifikasi data, bukan mengembalikan sebuah nilai. Di Go, prosedur diimplementasikan sebagai fungsi tanpa tipe kembalian yang menerima parameter pointer untuk menyampaikan hasil keluar ke pemanggil.
   - Pemahaman ini penting agar programmer dapat memilih antara prosedur dan fungsi berdasarkan kebutuhan: fungsi untuk komputasi nilai tunggal, prosedur untuk tindakan dengan efek samping atau keluaran berganda.

2. **Mekanisme Pass-by-Reference dengan Pointer:**
   - Penggunaan operator `&` (address-of) saat memanggil dan `*` (dereference) saat mengakses di dalam prosedur adalah inti dari mekanisme *pass-by-reference* di Go. Tanpa pointer, semua parameter di Go bersifat *pass-by-value* (salinan), sehingga perubahan di dalam fungsi tidak akan memengaruhi variabel asli di pemanggil.
   - Pola ini tampak jelas pada prosedur `cariMurid`, `hitungHuruf`, `cariTahun`, dan `tertuaTermuda`, di mana seluruh keluaran dikirimkan kembali melalui parameter pointer alih-alih melalui `return`.

3. **Pengelolaan Struct sebagai Parameter Pointer:**
   - Ketika prosedur perlu memodifikasi atau mengisi sebuah struct (seperti `penduduk` pada `tertuaTermuda`), meneruskan pointer ke struct (`*penduduk`) jauh lebih efisien dibanding menyalin seluruh isi struct. Ini mencegah penyalinan data yang tidak perlu (*unnecessary copy*) dan sekaligus memungkinkan prosedur menulis langsung ke memori struct aslinya.
