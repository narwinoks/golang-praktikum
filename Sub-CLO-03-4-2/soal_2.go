package main

import "fmt"

func main() {
	var x, y, z int
	var jumlahPertemuan int = 0
	fmt.Scan(&x, &y, &z)
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y == 0 && hari%z == 0 {
			jumlahPertemuan++
		}
	}

	fmt.Println(jumlahPertemuan)
}
