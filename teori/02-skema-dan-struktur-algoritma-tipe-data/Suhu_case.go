package main

import "fmt"

func main () {
	var umur int
	var suhu float32

	suhu = 37.6
	umur = 18

	fmt.Println("umur: ", umur)
	fmt.Println("suhu: ", suhu)
	fmt.Println("alamat memori dari var suhu: ", &suhu)
	fmt.Println("alamat memori dari var umur: ", &umur)
}