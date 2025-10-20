package main

import "fmt"

func main() {
	var hargaPergelas, jumlahGelas, serviceChange int
	var pajak float64
	fmt.Print("Masukan harga pergelas,jumlah gelas,pajak dan service change: ")
	fmt.Scanln(&hargaPergelas, &jumlahGelas, &pajak, &serviceChange)

	var subTotal int = hargaPergelas * jumlahGelas
	var totalPajak float64 = float64(subTotal) * (float64(pajak) / 100.0)
	var total float64 = float64(subTotal) + totalPajak + float64(serviceChange)
	fmt.Print("Sub Total ", total)

}
