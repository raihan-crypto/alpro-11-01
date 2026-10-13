package main

import "fmt"

func main(){
	var xi, total int
	fmt.Scan(&xi)
	for{
		total = total + xi
		fmt.Scan(&xi)
		if xi%2 !=0{
			break
		}
	}
	fmt.Println(total)
}