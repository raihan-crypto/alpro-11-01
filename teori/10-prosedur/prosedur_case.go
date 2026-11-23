package main 

import "fmt"


func add(a, b int) int {
	return a + b
}

func eat() {
	fmt.Println("saya lapar")
	fmt.Println("saya mengambil piring")
	fmt.Println("saya mengambil nasi dan beberapa lauk")
	fmt.Println("saya makan nasi dan lauk")
}

func hitungF(x, y int) {
	var hasil int
	hasil = 2*x + y - 5
	fmt.Println(hasil)
}

func hitungFF(x, y int) int {
	var hasil int
	hasil = 2*x + y - 5
	return hasil
}

func main () {
	fmt.Println (hitungFF(2, 3))
	x := 2
	y := 3
	hitungF(2, 3)
	hitungF(x, y)
	hitungF(x, 3)
	fmt.Println(add(2, 3))
	eat()
}