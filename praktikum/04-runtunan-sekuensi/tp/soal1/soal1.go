package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	fmt.Printf("1.  intNum > 5 → %v\n", intNum > 5)

	fmt.Printf("2.  intNum >= 5 && intOther < 11 → %v\n", intNum >= 5 && intOther < 11)

	fmt.Printf("3.  sngNum != -1 || intOther < 0 → %v\n", sngNum != -1 || intOther < 0)

	fmt.Printf("4.  !(intNum > 3) || intNum <= 5 → %v\n", !(intNum > 3) || intNum <= 5)

	fmt.Printf("5.  !(intOther >= intNum) → %v\n", !(intOther >= intNum))

	fmt.Printf("6.  0 - sngNum > 0 → %v\n", 0-sngNum > 0)

	fmt.Printf("7.  4 / 2 == intOther / intNum → %v\n", 4/2 == intOther/intNum)

	fmt.Printf("8.  intOther %% 2 == 0 → %v\n", intOther%2 == 0)

	fmt.Printf("9.  intOther + 2 * intNum != 30 || !(sngNum > 0) → %v\n", intOther+2*intNum != 30 || !(sngNum > 0))

	fmt.Printf("10. intOther > 0 && intNum > 0 || sngNum > 0 → %v\n", intOther > 0 && intNum > 0 || sngNum > 0)

	fmt.Printf("11. sngNum > 0 || (intNum >= 0 && -1 * intOther == -10) → %v\n", sngNum > 0 || (intNum >= 0 && -1*intOther == -10))

	fmt.Printf("12. intNum == 5 → %v\n", intNum == 5)

	fmt.Printf("13. intNum > 0 || (sngNum <= 0 && intOther == 13) → %v\n", intNum > 0 || (sngNum <= 0 && intOther == 13))

	fmt.Printf("14. !(!(!(!(intNum > 0)))) → %v\n", !(!(!(!(intNum > 0)))))
}
