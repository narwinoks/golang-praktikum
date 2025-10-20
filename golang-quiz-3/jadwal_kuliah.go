package main

import "fmt"

func toMinutes(jam, menit int) int {
	return (jam * 60) + menit
}
func cekBentrok(mulai1, selesai1, mulai2, selesai2 int) bool {
	return mulai1 < selesai2 && selesai1 > mulai2
}
func main() {
	var jm1, mm1, js1, ms1 int
	var jm2, mm2, js2, ms2 int
	var jm3, mm3, js3, ms3 int
	var mulai1, mulai2, mulai3 int
	var selesai1, selesai2, selesai3 int
	var adaBentrok bool

	fmt.Println("Masukan Jam Kuliah mapel 1 :")
	fmt.Scanln(&jm1, &mm1, &js1, &ms1)
	fmt.Println("Masukan Jam Kuliah mapel 2 :")
	fmt.Scanln(&jm2, &mm2, &js2, &ms2)
	fmt.Println("Masukan Jam Kuliah mapel 3 :")
	fmt.Scanln(&jm3, &mm3, &js3, &ms3)

	mulai1 = toMinutes(jm1, mm1)
	selesai1 = toMinutes(js1, ms1)

	mulai2 = toMinutes(jm2, mm2)
	selesai2 = toMinutes(js2, ms2)

	mulai3 = toMinutes(jm3, mm3)
	selesai3 = toMinutes(js3, ms3)

	adaBentrok = cekBentrok(mulai1, selesai1, mulai2, selesai2) ||
		cekBentrok(mulai1, selesai1, mulai3, selesai3) ||
		cekBentrok(mulai2, selesai2, mulai3, selesai3)
	fmt.Println("Jawal Kuliah Ada Bentrok ? ", adaBentrok)
}
