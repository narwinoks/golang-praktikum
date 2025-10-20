package main

import "fmt"

func convertTime(totalDetik int) string {
	var jam, sisaDetik, menit, detik int
	jam = totalDetik / 3600

	sisaDetik = totalDetik % 3600

	menit = sisaDetik / 60

	detik = sisaDetik % 60

	return fmt.Sprintf("%02d:%02d:%02d", jam, menit, detik)
}

func main() {
	var detikInput int
	var waktuFormatted string

	fmt.Print("Masukkan total detik (integer): ")
	fmt.Scan(&detikInput)

	waktuFormatted = convertTime(detikInput)
	fmt.Println("Hasil konversi:", waktuFormatted)
}
