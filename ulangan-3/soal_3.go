package main

import "fmt"

func main() {
	var totalDetik, jam, menit, detik int

	fmt.Scanln(&totalDetik)

	jam = totalDetik / 3600

	menit = (totalDetik / 60) % 60

	detik = totalDetik % 60

	fmt.Println(jam, menit, detik)
}
