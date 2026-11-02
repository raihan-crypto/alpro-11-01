```ruby
ALGORITMA Sistem_Perhitungan_Biaya_Laundry
DEKLARASI
jumlahPelanggan <- 0
totalBerat <- 0
totalPendapatan <- 0
DESKRIPSI
ULANGI
INPUT namaPelanggan
INPUT beratCucian
SELAMA beratCucian <= 0
TAMPILKAN "Berat harus lebih dari 0 kg"
INPUT beratCucian
AKHIR SELAMA
TAMPILKAN "Pilih layanan: 1. Reguler, 2. Ekspres"
INPUT pilihanLayanan
JIKA pilihanLayanan = 1 MAKA
tarif <- 5000
namaLayanan <- "Reguler"
JIKA TIDAK, JIKA pilihanLayanan = 2 MAKA
tarif <- 8000
namaLayanan <- "Ekspres"
JIKA TIDAK
TAMPILKAN "Pilihan layanan tidak valid"
KEMBALI KE INPUT PILIHAN LAYANAN
AKHIR JIKA
biayaAwal <- beratCucian * tarif
JIKA biayaAwal >= 50000 MAKA
diskon <- biayaAwal * 10 / 100
JIKA TIDAK
diskon <- 0
AKHIR JIKA
totalBayar <- biayaAwal - diskon
TAMPILKAN namaPelanggan, namaLayanan, beratCucian
TAMPILKAN biayaAwal, diskon, totalBayar
jumlahPelanggan <- jumlahPelanggan + 1
totalBerat <- totalBerat + beratCucian
totalPendapatan <- totalPendapatan + totalBayar
TANYAKAN "Proses pelanggan berikutnya? (ya/tidak)"
INPUT jawaban
SAMPAI jawaban = "tidak"
TAMPILKAN jumlahPelanggan
TAMPILKAN totalBerat
TAMPILKAN totalPendapatan
SELESAI
```