package main

import "fmt"

func main() {
	var n, i int
	var counter int = 0
	var bilangan int
	fmt.Println("Masukana n jumlah : ")
	fmt.Scanln(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&bilangan)
		counter = counter + (1 - (bilangan % 2))
	}
	println(counter)
}
