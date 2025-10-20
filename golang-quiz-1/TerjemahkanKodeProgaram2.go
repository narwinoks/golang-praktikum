package main

import "fmt"

func main() {
	var input int
	var val1, val2, val3, val4 int
	fmt.Print("Masukkan sebuah bilangan 4-digit: ")
	fmt.Scan(&input)

	val1 = input / 1000
	val2 = (input / 100) % 10
	val3 = (input / 10) % 10
	val4 = input % 10

	fmt.Println(val1*11, val2*11, val3*11, val4*11)
}
