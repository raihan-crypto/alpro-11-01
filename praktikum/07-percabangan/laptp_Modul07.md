# <h1 align="center">Tugas Pendahuluan Modul 07 - Percabangan</h1>
<p align="center">Raihan Ahmad Rahmadhani - 109092600025</p>

### 1. Perbedaan if, if-else, else if, dan switch serta Perbandingan switch Go vs Java

#### Deskripsi
Percabangan (*branching/selection*) merupakan salah satu struktur kontrol paling penting dalam algoritma dan pemrograman yang memungkinkan suatu program mengambil keputusan dan menentukan blok instruksi mana yang akan dieksekusi berdasarkan evaluasi kondisi logika boolean (`true` atau `false`).

Berikut adalah penjelasan perbedaan masing-masing struktur percabangan pada bahasa pemrograman Go:

1. **`if` (Percabangan Tunggal / Satu Arah):**
   - Struktur ini hanya mengevaluasi **satu kondisi**.
   - Blok kode di dalam kurung kurawal hanya akan dieksekusi apabila kondisi bernilai `true`.
   - Jika kondisi bernilai `false`, program tidak melakukan apa-apa pada blok tersebut dan langsung melanjutkan eksekusi ke baris kode setelah kurung kurawal tutup.
   - **Sintaks:**
     ```go
     if kondisi {
         // dieksekusi hanya jika kondisi bernilai true
     }
     ```

2. **`if-else` (Percabangan Biner / Dua Arah):**
   - Struktur ini menangani **dua kemungkinan alternatif keputusan**.
   - Apabila kondisi pada `if` bernilai `true`, maka blok di dalam `if` yang dijalankan.
   - Namun jika kondisi tersebut bernilai `false`, program akan secara otomatis mengeksekusi blok kode alternatif yang ada di dalam `else`.
   - **Sintaks:**
     ```go
     if kondisi {
         // dieksekusi jika kondisi bernilai true
     } else {
         // dieksekusi jika kondisi bernilai false
     }
     ```

3. **`else if` (Percabangan Majemuk / Bertingkat):**
   - Digunakan ketika terdapat **lebih dari dua kondisi alternatif** yang harus diuji secara berurutan (*cascade*).
   - Program mengevaluasi kondisi pertama pada `if`. Jika `false`, evaluasi beralih ke kondisi pada `else if` berikutnya, dan seterusnya.
   - Begitu ditemukan salah satu kondisi yang bernilai `true`, blok instruksi terkait akan langsung dieksekusi dan seluruh cabang `else if` maupun `else` di bawahnya akan diabaikan. Blok `else` opsional di akhir hanya akan dieksekusi jika seluruh kondisi di atasnya bernilai `false`.
   - **Sintaks:**
     ```go
     if kondisi1 {
         // dieksekusi jika kondisi1 true
     } else if kondisi2 {
         // dieksekusi jika kondisi1 false dan kondisi2 true
     } else {
         // dieksekusi jika seluruh kondisi di atas false
     }
     ```

4. **`switch` (Percabangan Seleksi Multi-Kondisi):**
   - Struktur yang dirancang untuk menyederhanakan pengujian nilai variabel terhadap sekumpulan kemungkinan konstanta atau ekspresi (`case`) secara lebih bersih, terstruktur, dan mudah dibaca dibanding menulis rantai `else if` yang panjang.
   - Di Go, pernyataan `switch` sangat fleksibel karena dapat mengevaluasi suatu nilai variabel tertentu (*expression switch*) maupun dijalankan tanpa ekspresi (*tagless switch*), di mana masing-masing `case` dapat memuat ekspresi boolean tersendiri.
   - **Sintaks:**
     ```go
     switch variabel {
     case nilai1:
         // aksi untuk nilai1
     case nilai2:
         // aksi untuk nilai2
     default:
         // aksi jika tidak ada case yang cocok
     }
     ```

---

#### Dua Perbedaan switch pada Go Dibandingkan dengan switch pada Java

1. **Perilaku *Implicit Break* (Otomatis Berhenti) vs *Fall-through*:**
   - **Pada Java:** Secara default, struktur `switch` di Java menerapkan perilaku *fall-through*. Artinya, setelah sebuah `case` yang cocok dieksekusi, aliran program akan terus merembes mengeksekusi baris perintah pada `case` berikutnya di bawahnya, kecuali jika programmer secara manual dan eksplisit menambahkan instruksi `break;` di akhir setiap `case`. Lupa menuliskan `break;` sering kali menjadi sumber kesalahan logika (*bug*) di Java.
   - **Pada Go:** Secara default, setiap blok `case` di Go memiliki sifat *implicit break* (otomatis keluar). Ketika suatu `case` bernilai cocok dan selesai dijalankan, eksekusi `switch` langsung selesai tanpa memerlukan kata kunci `break`. Sebaliknya, jika programmer secara khusus ingin melanjutkan eksekusi ke `case` di bawahnya, Go menyediakan kata kunci `fallthrough` yang harus ditulis secara eksplisit.

2. **Dukungan *Condition-less / Tagless Switch* (Switch Tanpa Ekspresi):**
   - **Pada Java:** Pernyataan `switch` di Java wajib mengevaluasi suatu variabel atau ekspresi spesifik yang diletakkan di dalam tanda kurung setelah kata kunci `switch` (misalnya `switch(angka)`), dan setiap `case` harus berupa nilai konstan yang bertipe sama (atau *pattern matching* pada Java versi modern). Java tidak memiliki sintaks `switch` kosong tanpa variabel pembanding.
   - **Pada Go:** Go mengizinkan penulisan `switch` tanpa ekspresi atau variabel di kepalanya (`switch { ... }`). Dalam bentuk ini, setiap baris `case` dapat berupa ekspresi logika boolean yang berdiri sendiri (seperti `case x > 0 && y < 10:`). Fitur ini menjadikan `switch` di Go sangat kuat sebagai pengganti rantai `if - else if - else` yang jauh lebih bersih dan idiomatis (*clean code*).

---

### 2. bilangan.go

```go
package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan == 0 {
		fmt.Println("Nol")
	} else {
		var paritas, tanda string

		if bilangan%2 == 0 {
			paritas = "Genap"
		} else {
			paritas = "Ganjil"
		}

		if bilangan > 0 {
			tanda = "Positif"
		} else {
			tanda = "Negatif"
		}

		fmt.Printf("%s, %s\n", paritas, tanda)
	}
}
```

##### Output

```text
Masukan: 8
Keluaran: Genap, Positif

Masukan: -3
Keluaran: Ganjil, Negatif

Masukan: 0
Keluaran: Nol
```

![Screenshot Output TP Soal 2](tp/soal%202/output.png)

#### Deskripsi
Program `bilangan.go` ini dibuat untuk menerima sebuah input bilangan bulat dari pengguna, kemudian menentukan apakah bilangan tersebut tergolong genap atau ganjil serta positif atau negatif. Khusus untuk angka nol (`0`), program memberikan perlakuan khusus karena nol bersifat netral (bukan positif dan bukan negatif).

Cara kerja program:
1. Program mendeklarasikan variabel `bilangan` bertipe `int` dan membaca input dari terminal menggunakan `fmt.Scan(&bilangan)`.
2. Pertama, program mengecek kondisi khusus angka nol menggunakan `if bilangan == 0`. Jika input adalah 0, program langsung mencetak `"Nol"` sesuai dengan spesifikasi kolom keluaran pada tabel soal.
3. Jika bilangan bukan nol (`else`), program melanjutkan ke evaluasi karakteristik bilangan menggunakan dua variabel penampung teks: `paritas` dan `tanda`.
4. Untuk menentukan paritas (genap atau ganjil):
   - Program menguji kondisi `bilangan % 2 == 0`.
   - Jika sisa bagi dengan 2 adalah 0, maka `paritas` diisi `"Genap"` (misalnya `8 % 2 == 0`).
   - Jika sisa bagi bukan 0, maka `paritas` diisi `"Ganjil"`. Di bahasa Go, operasi modulo pada bilangan negatif seperti `-3 % 2` menghasilkan `-1` (bukan 0), sehingga kondisi `else` secara otomatis dan akurat menangani bilangan ganjil positif maupun ganjil negatif.
5. Untuk menentukan tanda (positif atau negatif):
   - Program menguji kondisi `bilangan > 0`.
   - Jika lebih dari nol, `tanda` diisi `"Positif"`.
   - Jika kurang dari nol, `tanda` diisi `"Negatif"`.
6. Terakhir, program mencetak hasil gabungan paritas dan tanda dengan format string terpisah koma dan spasi: `fmt.Printf("%s, %s\n", paritas, tanda)`.

---

### 3. hari.go

```go
package main

import "fmt"

func main() {
	var nomorHari int
	fmt.Scan(&nomorHari)

	switch nomorHari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu (akhir pekan)")
	case 7:
		fmt.Println("Minggu (akhir pekan)")
	default:
		fmt.Println("Nomor hari tidak valid")
	}
}
```

##### Output

```text
Masukan: 3
Keluaran: Rabu

Masukan: 7
Keluaran: Minggu (akhir pekan)

Masukan: 9
Keluaran: Nomor hari tidak valid
```

![Screenshot Output TP Soal 3](tp/soal%203/output.png)

#### Deskripsi
Program `hari.go` ini dibuat menggunakan struktur percabangan `switch-case` untuk menerima input nomor hari antara 1 sampai 7, lalu menampilkan nama hari yang bersesuaian dengan ketentuan:
- Angka `1` adalah Senin, `2` adalah Selasa, `3` adalah Rabu, `4` adalah Kamis, dan `5` adalah Jumat.
- Angka `6` (Sabtu) dan `7` (Minggu) wajib menyertakan keterangan tambahan `(akhir pekan)`.
- Nomor di luar rentang 1–7 harus menampilkan pesan validasi `"Nomor hari tidak valid"`.

Cara kerja program:
1. Program membaca nomor hari dari masukan pengguna ke dalam variabel `nomorHari` dengan `fmt.Scan(&nomorHari)`.
2. Struktur `switch nomorHari` mengevaluasi nilai yang dimasukkan dan mencocokkannya ke cabang `case` yang sesuai:
   - Jika pengguna memasukkan angka 1 sampai 5, program mencetak nama hari kerja terkait (`"Senin"`, `"Selasa"`, `"Rabu"`, `"Kamis"`, atau `"Jumat"`).
   - Jika pengguna memasukkan angka 6 atau 7, program mencetak nama hari beserta keterangan akhir pekan, yaitu `"Sabtu (akhir pekan)"` dan `"Minggu (akhir pekan)"` sesuai instruksi pada soal.
   - Jika pengguna memasukkan angka selain 1 hingga 7 (seperti angka 9, 0, atau angka negatif), aliran eksekusi otomatis jatuh ke blok `default` yang mencetak pesan `"Nomor hari tidak valid"`.
3. Karena bahasa Go memiliki sifat *implicit break*, setelah sebuah `case` selesai dicetak program langsung keluar dari blok `switch` dengan aman tanpa merembes ke case lainnya.

---

### 4. terbesar.go

```go
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var terbesar int
	if a >= b && a >= c {
		terbesar = a
	} else if b >= a && b >= c {
		terbesar = b
	} else {
		terbesar = c
	}

	fmt.Println(terbesar)
}
```

##### Output

```text
Masukan: 4 9 2
Keluaran: 9

Masukan: 7 7 3
Keluaran: 7

Masukan: -5 -2 -8
Keluaran: -2
```

![Screenshot Output TP Soal 4](tp/soal%204/output.png)

#### Deskripsi
Program `terbesar.go` ini dibuat untuk mencari dan menentukan nilai bilangan terbesar dari tiga buah bilangan bulat yang dimasukkan oleh pengguna (`a`, `b`, dan `c`). Program dibangun murni menggunakan percabangan `if-else` bertingkat tanpa perulangan (`for`) dan tanpa fungsi bawaan `max` (seperti `math.Max`).

Cara kerja program:
1. Program mendeklarasikan tiga variabel bilangan bulat `a`, `b`, dan `c` bertipe `int`, kemudian membacanya sekaligus melalui `fmt.Scan(&a, &b, &c)`.
2. Program menyiapkan variabel `terbesar` untuk menyimpan hasil akhir nilai maksimum.
3. Logika perbandingan dilakukan menggunakan percabangan `if - else if - else` dengan operator relasional `>=` dan operator logika `&&`:
   - **Kondisi 1 (`if a >= b && a >= c`):** Menguji apakah nilai `a` lebih besar atau sama dengan `b` sekaligus lebih besar atau sama dengan `c`. Jika ya, maka `a` adalah bilangan terbesar (`terbesar = a`).
   - **Kondisi 2 (`else if b >= a && b >= c`):** Apabila kondisi pertama tidak terpenuhi, program menguji apakah nilai `b` lebih besar atau sama dengan `a` dan `c`. Jika ya, maka `b` adalah bilangan terbesar (`terbesar = b`).
   - **Kondisi 3 (`else`):** Jika `a` dan `b` terbukti bukan yang terbesar, maka secara eliminasi matematis nilai `c` pasti merupakan bilangan yang terbesar (`terbesar = c`).
4. Penggunaan operator `>=` (bukan hanya `>`) sangat krusial agar logika tetap bekerja sempurna ketika terdapat dua atau tiga bilangan yang bernilai kembar (seperti pada contoh input `7 7 3` di mana nilai terbesar adalah `7`).
5. Logika perbandingan ini juga sepenuhnya berlaku pada bilangan negatif (misalnya `-5 -2 -8`, di mana `-2` lebih besar dari `-5` dan `-8`).
6. Setelah percabangan selesai dievaluasi, program mencetak nilai variabel `terbesar` ke terminal menggunakan `fmt.Println(terbesar)`.

---

## Kesimpulan

Dari pengerjaan Tugas Pendahuluan Modul 07 mengenai percabangan dalam bahasa pemrograman Go ini, beberapa poin penting yang dapat disimpulkan antara lain:

1. **Pemilihan Struktur Percabangan yang Tepat:**
   - Gunakan percabangan tunggal `if` ketika hanya ada satu kondisi yang membutuhkan tindakan khusus.
   - Gunakan `if-else` jika terdapat dua alternatif tindakan yang saling berlawanan.
   - Gunakan `else if` atau `switch` untuk skenario multi-kondisi. `switch` sangat disarankan apabila pengujian melibatkan perbandingan nilai diskret agar kode lebih ringkas dan terhindar dari penumpukan indentasi (*indentation hell*).

2. **Keunggulan Struktur `switch` pada Go:**
   - Bahasa Go mengimplementasikan *implicit break* secara default pada setiap `case`, menghilangkan kemungkinan terjadinya *fall-through bug* yang sering dialami programmer bahasa C/C++ maupun Java.
   - Dukungan terhadap *condition-less switch* (`switch { ... }`) di Go memberikan alternatif sintaks yang jauh lebih rapi dan bersih dibanding rantai `if - else if - else` panjang saat menangani banyak kondisi ekspresi boolean.

3. **Penanganan Kasus Tepi (*Edge Cases*):**
   - Dalam merancang logika percabangan, penting untuk memperhatikan kasus-kasus khusus seperti nilai nol (yang netral terhadap tanda positif/negatif), sisa bagi modulo pada bilangan bulat negatif, serta penggunaan operator perbandingan relasional yang tepat (misalnya `>=` alih-alih `>`) agar program mampu menangani masukan bilangan bernilai kembar atau negatif secara akurat.
