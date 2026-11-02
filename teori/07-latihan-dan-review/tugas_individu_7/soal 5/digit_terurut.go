package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if len(s) <= 1 {
		fmt.Println("tidak terurut")
		return
	}

	ascending := true
	descending := true

	for i := 0; i < len(s)-1; i++ {
		if s[i] >= s[i+1] {
			ascending = false
		}
		if s[i] <= s[i+1] {
			descending = false
		}
	}

	if ascending {
		fmt.Println("ascending")
	} else if descending {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}
}

