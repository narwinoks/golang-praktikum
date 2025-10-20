package main

import (
	"fmt"
)

func keliling(jari int) float64 {
	return 2 * 22.0 / 7.0 * float64(jari)
}
func luas(jari int) float64 {
	return 22.0 / 7 * float64(jari) * float64(jari)
}

func main() {
	var jari1 int = 1
	fmt.Println("KELILING JARI 1  :", keliling(jari1))
	fmt.Println("LUAS JARI 1  :", keliling(jari1))

	var jari2 int = 5
	fmt.Println("KELILING JARI 5  :", luas(jari2))
	fmt.Println("LUAS JARI 5  :", luas(jari2))
	var jari3 int = 10
	fmt.Println("KELILING JARI 10  :", keliling(jari3))
	fmt.Println("LUAS JARI 10  :", keliling(jari3))

}
