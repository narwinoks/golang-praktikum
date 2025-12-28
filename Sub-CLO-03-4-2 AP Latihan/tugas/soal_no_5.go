package main

import (
	"fmt"
)

func main() {
	var u1, u2, u3 int

	// Membaca masukan 3 bilangan bulat
	// Scan akan membaca input yang dipisahkan spasi atau baris baru
	fmt.Scan(&u1, &u2, &u3)

	// Memeriksa kondisi
	if u1%2 != 0 && u2%2 != 0 && u3%2 != 0 {
		// Jika u1, u2, dan u3 semuanya ganjil
		fmt.Println("ganjil")
	} else if u1%2 == 0 && u2%2 == 0 && u3%2 == 0 {
		// Jika u1, u2, dan u3 semuanya genap
		fmt.Println("genap")
	} else {
		// Jika tercampur atau tidak memenuhi syarat di atas
		fmt.Println("bukan ganjil atau genap")
	}
}
