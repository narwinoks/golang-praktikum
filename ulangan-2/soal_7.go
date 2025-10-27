package main

import (
	"fmt"
)

func main() {
	// Hapus 'var i' dari sini
	var result int

	// Lakukan inisialisasi 'i' di sini:
	for i := 2; i <= 50; i++ {
		result = result + i
	}
	fmt.Println(result)
}
