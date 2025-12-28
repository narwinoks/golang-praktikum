package main

import "fmt"

func main() {
	var u1, u2, u3 int
	// Membaca tiga bilangan bulat
	fmt.Scan(&u1, &u2, &u3)

	// Cek Aritmatika: Selisih harus sama
	if u2-u1 == u3-u2 {
		fmt.Println("aritmatika")
		// Cek Geometri: Rasio harus sama (gunakan perkalian silang agar akurat)
	} else if u2*u2 == u1*u3 {
		fmt.Println("geometri")
		// Bukan keduanya
	} else {
		fmt.Println("bukan aritmatika atau geometri")
	}
}
