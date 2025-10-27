package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var terurut bool = true

	var digitKanan int = n % 10

	n = n / 10

	for n > 0 {
		var digitKiri int = n % 10
		if digitKiri >= digitKanan {
			terurut = false
			break
		}
		digitKanan = digitKiri

		n = n / 10
	}

	fmt.Println(terurut)
}
