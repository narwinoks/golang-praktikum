package main

import "fmt"

func checkBillanganGenap(bilangan int) bool {
	if bilangan%2 == 0 {
		return true
	}
	return false
}
func main() {
	var hasil bool
	var billangan int
	fmt.Scanln(&billangan)
	hasil = checkBillanganGenap(billangan)
	fmt.Println(hasil)
}
