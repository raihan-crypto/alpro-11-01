package main 

import "fmt"

func main () {
	var huruf string
	fmt.Scan(&huruf)

	if (huruf >= "A"  && huruf <= "Z") ||
	(huruf >= "a"  && huruf <= "z" ){
		fmt.Println("Alphabet")
	} else {
		fmt.Println ("Bukan Alphabet")
	}
}