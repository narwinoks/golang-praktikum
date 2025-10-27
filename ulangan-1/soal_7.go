package main

import "fmt"

func main() {
	var totalDetik int

	fmt.Scanln(&totalDetik)

	var jam int = totalDetik / 3600

	var menit int = (totalDetik / 60) % 60

	var detik int = totalDetik % 60

	fmt.Println(jam, menit, detik)
}
