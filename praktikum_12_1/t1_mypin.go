package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var d1, d2, d3, d4 int
	var count int = 0
	var i int = 0

	for i < n {
		var token int
		fmt.Scan(&token)

		isThirdDigitOne := (token/100)%10 == 1
		isSecondDigitZero := (token/10)%10 == 0

		if isThirdDigitOne && isSecondDigitZero {
			count++
			digit := token % 10
			if count == 1 {
				d1 = digit
			} else if count == 2 {
				d2 = digit
			} else if count == 3 {
				d3 = digit
			} else if count == 4 {
				d4 = digit
			}
		}

		i++
	}
	if count == 4 {
		fmt.Printf("%d%d%d%d\n", d1, d2, d3, d4)
	} else {
		fmt.Println("TIDAK ADA")
	}
}
