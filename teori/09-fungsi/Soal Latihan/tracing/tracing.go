package main

import "fmt"

func tambah(x, y int) int {
	x = x + y
	return x
}

func main() {
	a := 4
	b := tambah(a, 2)
	fmt.Println(a, b) // 4 6

	a = tambah(b, 3)
	fmt.Println(a, b) // 9 6
}

/*
Tracing:
  a := 4           → a = 4
  b := tambah(4,2) → x=4+2=6, return 6  → b = 6, a masih 4
  fmt.Println(a,b) → 4 6

  a = tambah(6,3)  → x=6+3=9, return 9  → a = 9, b masih 6
  fmt.Println(a,b) → 9 6

Baris yang mengubah a:
  - "a := 4"         → a diinisialisasi menjadi 4
  - "a = tambah(b,3)"→ a diperbarui menjadi hasil tambah(6,3) = 9

Nilai baru a (9) berasal dari pemanggilan tambah(b, 3),
yaitu b (6) + 3 = 9.
*/
