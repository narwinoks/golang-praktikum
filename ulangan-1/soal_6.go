package main

import "fmt"

func hitungSuhuToFarhiet(suhu float64) float64 {
	return (9.0 / 5.0 * suhu) + 32
}
func main() {
	var suhu, hasil float64
	fmt.Scanln(&suhu)
	hasil = hitungSuhuToFarhiet(suhu)
	fmt.Println(hasil)
}
