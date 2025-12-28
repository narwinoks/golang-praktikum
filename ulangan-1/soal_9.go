package main

import "fmt"

func checkBillanganGanjil(bilangan int) bool {
	if bilangan%2 != 0 {
		return true
	}
	return false
}
func main() {
	var hasil bool
	var billangan int
	fmt.Scanln(&billangan)
	hasil = checkBillanganGanjil(billangan)
	fmt.Println(hasil)
}
