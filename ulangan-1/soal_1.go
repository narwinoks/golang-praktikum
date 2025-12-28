package main

import "fmt"

func checkBillangan(bilangan int) bool {
	return bilangan < 0 && bilangan%2 != 0
}
func main() {
	var bilangan int
	var hasil bool

	fmt.Scanln(&bilangan)
	hasil = checkBillangan(bilangan)
	fmt.Printf("hasil", hasil)
}
