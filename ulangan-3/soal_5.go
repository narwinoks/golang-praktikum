package main

import "fmt"

func hitungLuasPermukaan(jariJari float64) float64 {
	return 4 * 3.14 * (jariJari * jariJari)
}

func main() {
	var jariJari, luasPermukaan float64
	fmt.Scanln(&jariJari)
	luasPermukaan = hitungLuasPermukaan(jariJari)
	fmt.Println(luasPermukaan)
}
