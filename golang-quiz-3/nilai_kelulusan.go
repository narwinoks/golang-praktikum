package main

import "fmt"

func checkNilaiKelulusan(nilaiTugas int, nilaiKuis int, nilaiUjianTulis int) bool {
	var nilaiAkhir float32 = (float32(nilaiTugas) * 62.5 / 100) + (float32(nilaiKuis) * 25 / 100) + (float32(nilaiUjianTulis) * 12.5 / 100)
	return nilaiAkhir >= 40.01
}
func main() {
	var nilaiTugas, nilaiKuis, nilaiUjianTulis int
	fmt.Println("Masukan Nilai Tugas : ")
	fmt.Scanln(&nilaiTugas)
	fmt.Println("Masukan Nilai Kuis : ")
	fmt.Scanln(&nilaiKuis)
	fmt.Println("Masukan Nilai Kuis : ")
	fmt.Scanln(&nilaiUjianTulis)
	fmt.Println("Apakah Anda Lulus : ", checkNilaiKelulusan(nilaiTugas, nilaiKuis, nilaiUjianTulis))
}
