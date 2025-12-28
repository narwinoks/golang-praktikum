package main

import "fmt"

func main() {
	var n, i, bilanganGenap, hasil int
	var sign int = -1
	fmt.Scanln(&n)
	sign = sign * -1
	bilanganGenap = 1 * 2
	hasil = sign * bilanganGenap
	fmt.Println(hasil)
	for i = 2; i <= n; i++ {
		sign = sign * -1
		bilanganGenap = i * 2
		hasil = sign * bilanganGenap
		fmt.Println("")
		fmt.Println(hasil)
	}
}
