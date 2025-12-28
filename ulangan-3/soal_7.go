package main

import "fmt"

func main() {
	var jumlah int = 0
	var bilangan int = 2
	for bilangan <= 50 {
		jumlah = jumlah + bilangan
		bilangan++
	}
	fmt.Println(jumlah)
}
