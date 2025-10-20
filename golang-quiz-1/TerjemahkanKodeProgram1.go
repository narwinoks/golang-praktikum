package main

import "fmt"

func main() {
	var val1 int
	var val2, val3 float64

	fmt.Print("Masukkan nilai untuk val1 (integer): ")
	fmt.Scan(&val1)

	val2 = (float64(val1) * 9.0 / 5.0) + 32

	val3 = float64(val1) + 273.15

	fmt.Printf("Hasil val2: %f\n", val2)
	fmt.Printf("Hasil val3: %f\n", val3)
}
