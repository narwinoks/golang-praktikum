package main

import "fmt"

func main() {
	var bilangan, jumlahGanjil int

	// Kita melakukan perulangan 5 kali untuk membaca 5 input
	for i := 0; i < 5; i++ {
		fmt.Scan(&bilangan)
		// Cek apakah bilangan tersebut ganjil
		if bilangan%2 != 0 {
			jumlahGanjil++
		}
	}

	// Menentukan output berdasarkan jumlah bilangan ganjil yang ditemukan
	if jumlahGanjil == 5 {
		fmt.Println("ganjil semua")
	} else if jumlahGanjil == 0 {
		fmt.Println("tidak ada bilangan ganjil")
	} else {
		// Jika jumlahnya 1, 2, 3, atau 4
		fmt.Println("ganjil sebagian")
	}
}
