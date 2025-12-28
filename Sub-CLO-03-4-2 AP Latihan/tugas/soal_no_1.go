package main

import (
	"fmt"
)

func main() {
	var rankingKelas int

	fmt.Print("Masukkan ranking kelas anak (bilangan asli): ")

	fmt.Scanln(&rankingKelas)
	tentukanHadiah(rankingKelas)
}

func tentukanHadiah(ranking int) string {
	if ranking > 0 && ranking <= 5 {
		return "mendapat hadiah"
	}
	return ""
}
